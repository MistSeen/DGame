package tcp

import ("net"
  "time"
  "log"
	"sync/atomic"
	"sync"
)
const (
	ServerState_Init 		=iota
	ServerState_Listening
	ServerState_Close			//close: close listen and all connect
	ServerState_Shutdown			//server is shutdown, close listen  all connect  and dispose all resource
)

/**
 TcpServer is server
 //todo:
*/
type TcpServer struct {
	Addr string

	state  int
	ls    net.Listener
	conns TcpConnSet

	connCur int 			//当前链接数量
	connMax int			//最大连接数

	waits *sync.WaitGroup		//??这个是干什么的呢....
	connSucceFunc  ConnCallBackFunc
	connErrorFunc  ConnCallBackFunc
}

func NewTcpServer(addr string,maxConn int) *TcpServer{
	tcpServer := &TcpServer{Addr:addr,connMax:maxConn}
	tcpServer.waits =&sync.WaitGroup{}
	return tcpServer
}

//------------------------------------------------------------//
func (s *TcpServer) Init(succeFunc ConnCallBackFunc,errorFunc ConnCallBackFunc) {
	s.connSucceFunc = succeFunc;
	s.onConnectLost = errorFunc;
}
/*
StartAndListren
*/
func (s *TcpServer) StartAndListen() {


	//TODO::检查参数配置是否正确

	//监听端口
	if err := s.Listen(); err != nil {
		panic(err)
		return
	}
}
/*
Listren:一般用在Close状态后,重新Listen的时候
*/
func (s *TcpServer) Listen() error {
	s.waits.Add(1)
	defer s.waits.Done()

	addr, err := net.ResolveTCPAddr("tcp4", s.Addr)
	if err != nil {
		return err
	}
	ls, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err;
	}
	//TODO::如果有其他程序已经监听了这个端口,报错返回上层处理
	if !atomic.CompareAndSwapInt32(&s.state,ServerState_Init,ServerState_Listening) &&
		!atomic.CompareAndSwapInt32(&s.state,ServerState_Close,ServerState_Listening){
		return
	}
	s.ls = ls
	go s.goListening()      //accept 后,加入工作线程
	return nil
}
/*
Close is stop listerning and disconnect all connect
*/
func (s *TcpServer) Close() error{
	//1.状态判定,不能Close 2次
	if ! atomic.CompareAndSwapInt32(&s.state,ServerState_Listening,ServerState_Close){
		log.Print("server is closed...")
		return
	}
	if s.ls !=nil{
		s.ls.Close();
	}
	for c := range s.conns {
		c.Close()
	}
	atomic.StoreInt32(&s.connCur,0)
	return nil
}
func (s *TcpServer) Shutdown() {
	if atomic.LoadInt32(&s.state)==ServerState_Shutdown{
		return
	}
	atomic.StoreInt32(&s.state,ServerState_Shutdown)
	log.Fatal("server is shutdown...")
	if s.ls !=nil{
		s.ls.Close();
	}
	for c := range s.conns {
		c.Close()
	}
	s.waits.Wait()
	//TODO:: 清理资源
}

func (s* TcpServer) GetCurrConnCount() int {
	return s.connCur
}
func (s* TcpServer) GetMaxConnCount() int{
	return  s.connMax
}



//------------------------------------------------------------//


const tcpserver_accept_timeout time.Duration = 1000 * time.Millisecond

func (s *TcpServer) goListening(){
	s.waits.Add(1)
	defer s.waits.Done()

	var delay time.Duration
	for {
		//s.ls.(net.TCPListener).SetDeadline(time.Now().Add(1e9)) ?? 是否需要暂定监听
		conn, err := s.ls.Accept()
		//当accpet发生错误的的时候,超过1秒链接不上
		//1.说明系统繁忙,稍后再试一试
		//2.说明端口被占用/关闭,退出接受accpet状态
		if err != nil {
			if opErr,ok := err.(net.OpError);ok && opErr.Timeout(){
				continue
			} else if acceptErr, ok := err.(net.Error); ok && acceptErr.Temporary() {
				if delay == 0 {
					delay = 5 * time.Millisecond
				} else {
					delay *= 2
				}
				if delay > tcpserver_accept_timeout {
					delay = tcpserver_accept_timeout
				}

				log.Fatal("accept error: %v;  retry delay %d ms", acceptErr, delay) //TODO: log.warn xxxx
				time.Sleep(delay)
				continue
			}
			s.onAcceptErrorClose(err)
			return
		}
		delay = 0

		tcpConn := s.createConn(conn)
		go tcpConn.working()
		//TODO::放入缓存池,上锁
		s.conns[conn] = tcpConn
	}
}
func (s *TcpServer) onAcceptErrorClose(err error){
	log.Fatal("accept routine quit.error:%s", err.Error())		//TODO: log.err xxxx
	atomic.StoreInt32(*s.state,ServerState_Close)
	s.ls.Close()
	s.ls=nil
}
func (s *TcpServer) createConn(c net.Conn) *TcpConn{
	conn :=NewTcpConn(c,s.onConnectSuccess, s.onConnectLost);
	conn.SetReadDeadline(30 * time.Second) //
	//TODO:: 协议处理

	return conn
}
//------------------------------------------------------------//
func (s *TcpServer) onConnectSuccess(c *TcpConn,_ error) {
	s.incConCount(c)
	log.Printf("connect suncces is ok")
	go c.HeartHoop(30 * time.Second);	//心跳  TODO::所有链接心跳都纳入心跳定时器管理

	if s.connSucceFunc != nil {
		s.connSucceFunc(c, nil)
	}
}
func (s *TcpServer) onConnectLost(c *TcpConn,err error) {
	s.decConCount(c)
	//当链接丢失的时候,触发链接错误
	log.Fatalf("connect error ", err)
	if s.connErrorFunc != nil {
		s.connErrorFunc(c, err)
	}
}

func( s* TcpServer) incConCount(conn *TcpConn){
	if atomic.AddInt32(&s.connCur,1)+1> s.connMax{
		//TODO::通知上层,链接已经满了
	}
}
func( s* TcpServer) decConCount(conn *TcpConn){
	if atomic.AddInt32(&s.connCur,-1)< s.connMax{
		//TODO::通知上层,链接可用
	}
}