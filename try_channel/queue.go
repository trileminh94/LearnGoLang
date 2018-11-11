package queue

var message = make(chan string)

func send(mesg string) {
	go func() { message <- mesg }()
}

func receive() string {
	return <-message
}
