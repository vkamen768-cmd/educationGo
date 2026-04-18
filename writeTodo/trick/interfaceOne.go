package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	Email string
}

func (e *EmailNotifier) Notify(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.Email, message)
}

func main() {
	var n Notifier

	e := EmailNotifier{Email: "vlad@gmail.com"}
	e.Notify("Hello")
	n = &e
	n.Notify("fff")
}
