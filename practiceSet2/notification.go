package practiceset2

import "fmt"

type Notification interface {
	send(msg string)
}

type EmailNotifier struct {
	Email string
}

func (a EmailNotifier) send(msg string) {
	fmt.Printf("Send Email to %s:%s\n", a.Email, msg)
}

type Phone struct {
	Mobile string
}

func (b Phone) send(msg string) {
	fmt.Printf("Send SMS to %s:%s\n", b.Mobile, msg)
}

type PushNotifier struct {
	SMS string
}

func (c PushNotifier) send(msg string) {
	fmt.Printf("Send notification to %s:%s\n", c.SMS, msg)
}

func NotificationSystem() {

	email := EmailNotifier{Email: "abc@gmail.com"}
	phone := Phone{Mobile: "1234567890"}
	push := PushNotifier{SMS: "Device-001"}

	notifier := []Notification{email, phone, push}

	for _, n := range notifier {
		n.send("Hello, welcome to GOlang.")
	}
}
