package basics

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type CreditCard struct {
	CardNumber string
}

type UPI struct {
	ID string
}

type Cash struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Amount of %.2f is paid from[%s]\n", amount, c.CardNumber)
}
func (c UPI) Pay(amount float64) {
	fmt.Printf("Amount of %.2f is paid from[%s]\n", amount, c.ID)
}
func (c Cash) Pay(amount float64) {
	fmt.Printf("Amount of %.2f is paid Cash", amount)
}

func PaymentMethod() {
	cc := CreditCard{CardNumber: "1234-5678-9012"}
	upi := UPI{ID: "qwerty123"}
	cash := Cash{}

	Amt := []Payment{cc, upi, cash}

	for _, amt := range Amt {
		amt.Pay(250.34)
	}
}
