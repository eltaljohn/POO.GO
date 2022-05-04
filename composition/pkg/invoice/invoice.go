package invoice

import (
	"github.com/eltaljohn/composition/pkg/customer"
	"github.com/eltaljohn/composition/pkg/invoiceitem"
)

// Invoice is a struct of invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}

// New returns a new invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{country: country, city: city, client: client, items: items}
}
