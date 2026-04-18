package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	Email string
}
type SMSNotifier struct {
	Phone string
}

type User struct {
	Name      string
	Notifiers []Notifier
}

type ConsoleNotifier struct{}

func (e *EmailNotifier) Notify(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.Email, message)
}

func (s *SMSNotifier) Notify(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.Phone, message)
}

func (u User) Send(message string) {
	for _, v := range u.Notifiers {
		v.Notify(message)
	}
}

func (e *ConsoleNotifier) Notify(message string) {
	fmt.Printf("Console: %s\n", message)
}

func main() {

	notifiers := []Notifier{
		&EmailNotifier{Email: "Vlad@gmail.com"},
		&SMSNotifier{Phone: "12345678"},
		&ConsoleNotifier{},
	}

	u := User{Name: "Vlad", Notifiers: notifiers}
	u.Send("Hello Word!")

}
