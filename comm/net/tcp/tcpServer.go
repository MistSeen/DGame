package tcp

import (
    
)

type TcpServer{
    Conn TcpConn;
}


func TcpServer NewTcpServer(){

    server TcpServer* =&TcpServer{};

    net.TCPConn
    server.Conn = newTcpConn(conn net.TCPConn); 

}