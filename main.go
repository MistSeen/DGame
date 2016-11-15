// server project main.go
package main

import (
	"fmt"
	"net"
	"time"
)

var x chan int = make(chan int, 10)
var LS *net.TCPListener

func dd(t time.Duration, v int) {
	time.Sleep(t * time.Second)
	x <- v
	if v == 10 {
		//close(x)
		fmt.Println("close x ->")
		//LS.Close()
	}
	fmt.Printf("dd ->%v \n", v)
}

func ll() {
	addr, _ := net.ResolveTCPAddr("tcp4", ":9999")
	ls, _ := net.ListenTCP("tcp", addr)
	LS = ls
	fmt.Printf("v = %v \n", addr)

	//	for {
	//		fmt.Println("next accept")
	//		conn, err := ls.Accept()
	//		if err != nil {
	//			goto EndGo
	//		}
	//		conn.(*net.TCPConn).SetKeepAlive(true)
	//		conn.(*net.TCPConn).SetKeepAlivePeriod(2 * time.Second)
	//		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	//	}
	//EndGo:
	//	fmt.Println("end loop...")

}

func listen() {

	addr, _ := net.ResolveTCPAddr("tcp4", ":9999")
	ls, _ := net.ListenTCP("tcp", addr)
	fmt.Printf("v = %v \n", addr)

	for {
		select {
		case xx := <-x:

			fmt.Printf("v = %v \n", xx)

			if xx == 10 {
				fmt.Println("break")
				ls.Close()
				goto EndGo
			}
		default:
		}
		fmt.Println("next accept")
		ls.SetDeadline(time.Now().Add(1e9))
		conn, err := ls.Accept()

		if err != nil {
			if acceptErr, ok := err.(net.Error); ok {
				if acceptErr.Timeout() {
					fmt.Println("acceptErr....err:%v", acceptErr.Error())
					continue
				}

			}
			fmt.Println("connect....err:%v", err.Error())
			goto EndGo
		}
		//conn.SetKeepAlive(true)
		//conn.SetKeepAlivePeriod(d)
		conn.SetDeadline(time.Now().Add(2 * time.Second))
		//time.Sleep(1 * time.Second)
		fmt.Println("connect....ok")
	}
EndGo:
	fmt.Println("end loop...")
}
func main() {

	//1. 如果x 有消息进入  并且消息==1 , 跳出循环
	//2. 如果x 没有消息进入,阻塞等到消息
	//3.
	go dd(5, 1)
	go dd(10, 10)
	go listen()
	//go ll()
	for {
		time.Sleep(20 * time.Second)
	}
	fmt.Println("end World")
}
