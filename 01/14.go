package main

import "fmt"

type PaymentGateway interface {
	Pay(amont float64) string
}

type PayPal struct{}

func (PayPal) Pay(amount float64) string { return "Paid via PayPal" }

type Stripe struct{}

func (Stripe) Pay(amount float64) string { return "Paid via Stripe" }

type COD struct{}

func (COD) Pay(amount float64) string { return "Cash On Delivery selected" }

func Checkout(pg PaymentGateway) {
	fmt.Println(pg.Pay(500))
}

func main() {
	Checkout(PayPal{})
	Checkout(Stripe{})
	Checkout(COD{})
}
