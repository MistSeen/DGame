package tcp

import (
	"net"
	"time"
)

type TcpConnPool map[net.Conn]struct{}

type TcpConn struct {
	conn      net.TCPConn
	Chans	  chan []byte 
}

func newTcpConn(conn net.TCPConn) *TcpConn {
	tcpConn = &TcpConn{conn: conn}
	tcpConn.Chans = make(chan []byte, 10]) 
	tcpConn.SetReadDeadline(60*1000);
	go func ()  {
		for b := range tcpConn.Chans {
			if b == nil {
				break
			}

			_, err := conn.Write(b)
			if err != nil {
				break
			}
		}

		conn.Close()
		tcpConn.Lock()
		tcpConn.closeFlag = true
		tcpConn.Unlock()
	}
	return tcpConn
}

func (c *TcpConn) SetReadDeadline(millisecond int) (err error) {
	c.conn.SetReadDeadline(time.Now().Add(millisecond/time.Millisecond) )
} 

func (c *TcpConn) Close() (err error) {
	c.conn.Close()
}

func (c *TcpConn) Disponse() (err error) {

}

func (c *TcpConn) Recv() (ret interface{}, err error) {

}

func (c *TcpConn) Send(data interface{}) (err error) {

}
