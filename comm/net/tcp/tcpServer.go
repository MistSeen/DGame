package tcp

import ("net"
  "time"
  "log" 
)

type TcpServer struct {
	Addr string

	ls    net.Listener
	conns TcpConnSet

	connCallBack  ConnCallBackFunc
	errorCallBack ConnCallBackFunc
}

func (s *TcpServer) Start() {
	//检查参数配置是否正确
	s.init()

	//监听端口
	if err := s.listen(); err != nil {
		panic(err)
		return
	}
	//accept 后,加入工作线程
	s.working()
}

func (s *TcpServer) Stop() error {

	s.ls.Close()
	for c := range s.conns {
		c.Close()
	}

}
func (s *TcpServer) init() error {
	return nil
}
func (s *TcpServer) listen() error {

	addr, _err := net.ResolveTCPAddr("tcp4", s.Addr)
	if ls, err := net.ListenTCP("tcp", addr); err != nil {
		panic(err)
	}
	
	//TODO::如果有其他程序已经监听了这个端口,log....
	
	s.ls = ls
}
func (s *TcpServer) working() error {

	var delay time.Duration
	var max time.Duration = 1000 * time.Millisecond
	for {
		conn, err := s.ls.Accept()
		//链接的时候,超过1秒链接不上
		//说明系统繁忙,稍后再试一试
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if delay == 0 {
					delay = 5 * time.Millisecond
				} else {
					delay *= 2
				}
				if delay > max {
					delay = max
				}
				log.Fatal("accept error: %v; retrying in %v", err, delay)
				time.Sleep(delay)
				continue
			}
			return
		}
		delay = 0

		tcpConn := NewTcpConn(conn)

		tcpConn.SetDeadline(30) //

		tcpConn.SetCallBack(
		func(net.Conn,err error){
			log.Printf("connect suncces is ok") 
		},
		func(net.Conn,err error){
			log.Fatalf("connect error ",err) 	
		})
		
		//TODO::放入缓存池,上锁
		s.conns[conn] = struct{}{}

	}

}
