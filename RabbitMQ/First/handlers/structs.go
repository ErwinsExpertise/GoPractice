package handlers

type Queue struct {
	Messages []Message
}

type Message struct {
	OPCode uint8
}

var Listener Queue
