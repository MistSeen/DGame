package net

type Response interface {
	Conn
	OnStrarting()
	OnDoing()
	OnCloseing()
}
