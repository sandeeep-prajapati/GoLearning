package main

import "fmt"

type Listener interface {
	Update(data string)
}
type EventBus struct {
	listeners []Listener
}

func (bus *EventBus) Register(l Listener) {
	bus.listeners = append(bus.listeners, l)
}

func (bus *EventBus) Broadcast(msg string) {
	for _, l := range bus.listeners {
		l.Update(msg)
	}
}

type EmailListener struct{}

func (EmailListener) Update(data string) {
	fmt.Println("Email:", data)
}

type SMSListener struct{}

func (SMSListener) Update(data string) {
	fmt.Println("SMS:", data)
}

func main() {
	bus := EventBus{}
	bus.Register(EmailListener{})
	bus.Register(SMSListener{})

	bus.Broadcast("New User Registered!")
}
