package netting

type Request interface {
	Conn
	OnStrarting()
	OnDoing()
	OnCloseing()
}
