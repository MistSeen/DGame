package net

type Request interface {
	Conn
	OnStrarting()
	OnDoing()
	OnCloseing()
}
