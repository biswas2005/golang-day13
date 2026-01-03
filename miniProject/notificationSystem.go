package miniproject

import "fmt"

type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Send(msg string) {
	fmt.Printf("Sending Email to %s:%s\n", e.Email, msg)
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(msg string) {
	fmt.Printf("Sending SMS to %s:%s\n", s.Phone, msg)
}

type WatsappNotifier struct {
	Number string
}

func (w WatsappNotifier) Send(msg string) {
	fmt.Printf("Sending Message to %s:%s\n", w.Number, msg)
}

func UserNotifier(n Notifier, message string) {
	n.Send(message)
}

func System() {
	email := EmailNotifier{Email: "xyz12@gmail.com"}
	sms := SMSNotifier{Phone: "0987654321"}
	watsapp := WatsappNotifier{Number: "1234567890"}

	UserNotifier(email, "Welcome to Google")
	UserNotifier(sms, "Your OTP is 1209")
	UserNotifier(watsapp, "Your order has been Shipped.")
}
