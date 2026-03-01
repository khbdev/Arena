package domain

type MessageProcessor interface {
	ProcessMessage(body []byte)
}