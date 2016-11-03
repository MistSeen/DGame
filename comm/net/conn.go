package net
 
type OnConnettingFunc func(err error);


type Conn interface {
	Recv() interface{} 	error
	Send(interface{}) 	error 
	Close()				error

    connCallBack	 OnConnettingFunc; 
	errorCallBack	 OCloseingFunc; 
}
