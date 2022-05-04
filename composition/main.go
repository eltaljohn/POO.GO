package main

import (
	"fmt"
	"github.com/eltaljohn/composition/pkg/customer"
	"github.com/eltaljohn/composition/pkg/invoice"
	"github.com/eltaljohn/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Bogotá",
		customer.New(
			"John",
			"Calle falsa 123",
			"123423424234",
		),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso Go", 12.34),
			invoiceitem.New(1, "Curso de POO con Go", 43.34),
			invoiceitem.New(1, "Curso de testing Go", 90.34),
		),

	)
	i.SetTotal()

	fmt.Printf("%+v", i)
}
