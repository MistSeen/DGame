package tcp

import (
	"errors"
	"log"
	"net"
	. "server/core/packet"
	"sync/atomic"
	"time"
)

const (
	ConnState_Working = iota
	ConnState_Close
)

type ConnCallBackFunc func(*TcpConn, error)

type TcpConnSet map[*TcpConn]struct{}

type TcpConnCfg struct {
	Unpacker     IUnpacker //解包接口
	BuffRecvSize int       //接受Buff大小
	BuffSendSize int       //发送Buff大小

	HeartLiveMax   int              //最大心跳次数
	SendTimeoutSec int              //发送超时
	ConnSuccFunc   ConnCallBackFunc //链接成功后,回调
	ConnLostFunc   ConnCallBackFunc //链接丢失后,回调

}

type TcpConn struct {
	id uint64 //connect Id

	Chans     chan []byte
	sendChan  chan interface{} //write chan
	heartLive int              //心跳次数

	state int         //链接状态 ConnState_XXXX
	conn  net.TCPConn //原生链接

	localAddr  string
	remoteAddr string

	cfg TcpConnCfg //tcp的配置
}

var globalConnID uint64

func DefaultTcpConCfg() *TcpConnCfg {
	cfg := &TcpConnCfg{}
	cfg.BuffRecvSize = 1024
	cfg.BuffSendSize = 1024
	cfg.HeartLiveMax = 10
	cfg.SendTimeoutSec = 10
	return cfg
}

func NewTcpConn(c net.Conn, cfg *TcpConnCfg) *TcpConn {

	if cfg == nil {
		cfg = DefaultTcpConCfg()
	}

	conn := &TcpConn{id: atomic.AddUint64(&globalConnID, 1), conn: c}
	conn.localAddr = c.LocalAddr().String()
	conn.remoteAddr = c.RemoteAddr().String()

	if cfg.BuffSendSize > 0 {
		conn.sendChan = make(chan interface{}, cfg.BuffSendSize)
	}

	conn.Chans = make(chan []byte, 10)

	return conn
}

func (c *TcpConn) SetReadDeadline(d time.Duration) (err error) {
	c.conn.SetKeepAlive(true)
	c.conn.SetKeepAlivePeriod(d)
	return c.conn.SetReadDeadline(time.Now().Add(d))
}

func (c *TcpConn) Close() (err error) {
	if atomic.LoadInt32(&c.state) != ConnState_Working {
		return ConnErr_ConnIsNotWorking
	}
	c.conn.SetLinger(c.cfg.SendTimeoutSec)
	select {
	//等待
	case <-time.After(time.Duration(c.cfg.SendTimeoutSec)):
		c.closing(nil) //	timeout, close the connection
	}
	return nil
}

func (c *TcpConn) Disponse() {
	//SetLinger设定一个连接的关闭行为当该连接中仍有数据等待发送或者确认.
	//如果sec<0(默认形式),操作系统将在后台完成发送数据操作;
	//如果sec=0,操作系统将任何未发送或者未确认的数据丢弃;
	//如果sec>0,数据将在后台进行发送,这点和sec<0时效果一致.然而,在一些操作系统中,当sec秒之后,系统将任何未发送的数据丢弃.
	c.conn.SetLinger(0)
	c.closing(nil)
}

func (c *TcpConn) Recv() (ret interface{}, err error) {

	if atomic.LoadInt32(&c.state) != ConnState_Working {
		return ConnErr_ConnIsNotWorking
	}
	for {
		if c.cfg.Unpacker != nil {
			if packErr:= c.cfg.Unpacker.Unpacking(c,ret);packErr!=nil{
				err =packErr
				return
			}
		}else {
		   //TODO::直接读取
		}
	}
	return nil, errors.New("Not Implemented")
}

func (c *TcpConn) Send(data interface{}) (err error) {
	if atomic.LoadInt32(&c.state) != ConnState_Working {
		return ConnErr_ConnIsNotWorking
	}
	//TODO::
	return errors.New("Not Implemented")
}

func (c *TcpConn) HeartHoop(d time.Duration) {
	if c.cfg.HeartLiveMax <= 0 {
		return
	}
	time.AfterFunc(d, func() {
		if atomic.LoadInt32(*c.state) != ConnState_Working {
			return
		}
		if atomic.AddInt32(&c.heartLive, 1) > c.heartLiveMax {
			c.conn.SetReadDeadline(time.Now().Add(d))
		} else {
			c.closing(errors.New("Connect is HeartHoop Timeout"))
		}
	})
}
func (c *TcpConn) GetConnID() int64 {
	return c.id
}

// GetRemoteAddress return the remote address of the connection
func (c *TcpConn) GetRemoteAddress() string {
	return c.remoteAddr
}

// GetLocalAddress return the local address of the connection
func (c *TcpConn) GetLocalAddress() string {
	return c.localAddr
}
func (c *TcpConn) GetConnCfg() *TcpConnCfg {
	return c.cfg
}
func (c *TcpConn) working() {
	defer func() {
		e := recover()
		if e != nil {
			log.Fatalf("tcp:%v <-> %v Painc %v", c.remoteAddr, c.localAddr, e)
		}
		c.closing(e)
	}()
	atomic.StoreInt32(*c.state, ConnState_Working)

	c.onConnSucessing()
	go c.reading()
	go c.writing()
}
func (c *TcpConn) onConnSucessing() {
	if c.cfg.ConnSuccFunc != nil {
		c.cfg.ConnSuccFunc(c, nil)
	}
}
func (c *TcpConn) onConnLosting(err error) {
	if c.cfg.ConnLostFunc != nil {
		c.cfg.ConnLostFunc(c, err)
	}
}
func (c *TcpConn) reading() {

}
func (c *TcpConn) writing() {

}
func (c *TcpConn) closing(err error) {
	if !atomic.CompareAndSwapInt32(&c.state, ConnState_Working, ConnState_Close) {
		return ConnErr_ConnIsNotWorking
	}
	c.conn.Close()
	c.onConnLosting(err)
}
