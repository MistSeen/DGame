package net

import (
	"net"
)
 
type ConnCallBackFunc func(net.Conn,err error);


type Conn interface {
	Recv() interface{} ,	error
	Send(interface{}) 	error 
	Close()				error


}
