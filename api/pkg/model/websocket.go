package model

type IWsHandler interface {
	Send(msg string) error
	Receive() (string, error)
	Close() error
}
