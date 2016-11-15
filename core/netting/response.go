package netting

type Response interface {
	Conn
	OnStrarting()
	OnDoing()
	OnCloseing()
}
