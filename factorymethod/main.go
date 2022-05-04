package main

import "fmt"

type PaymentMethod interface {
	Pay()
}

type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado por Paypal")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado por efectivo")
}

type CreditCard struct{}

func (cr CreditCard) Pay() {
	fmt.Println("Pagado por tarjeta de crédito")
}

func Factory(method uint) PaymentMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digíte un número de método de pago")
	fmt.Println("\t 1:Paypal 2:Efectivo 3:Tarjeta de crédito")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Debe digitar un método válido")
	}

	if method > 3 {
		panic("Debe digitar un método válido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
