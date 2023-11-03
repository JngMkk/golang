package main

import koreapost "github.com/JngMkk/practice/interface2/koreaPost"

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &koreapost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
