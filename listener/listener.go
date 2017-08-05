package listener

type Listener interface {
	Listen()
}

func StartListening(listener Listener) {
	listener.Listen()
}
