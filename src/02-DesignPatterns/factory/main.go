package main

import (
	"fmt"
	"os"
)

//START1 OMIT
type PaymentMethod interface {
	Pay(amount float64) string
}

//END1 OMIT

//START2 OMIT
const (
	Cash      = 1
	Debitcard = 2
)

//END2 OMIT

//START3 OMIT
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case Debitcard:
		return new(DebitcardPM), nil
	default:
		return nil, fmt.Errorf("method no registered: %d", m)
	}
}

//END3 OMIT

//START4 OMIT
type CashPM struct{}
type DebitcardPM struct{}

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f payment made by cash", amount)
}

func (d *DebitcardPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f payment made by Debitcard", amount)
}

//END4 OMIT

//START5 OMIT
func main() {
	p, err := GetPaymentMethod(Cash)
	if err != nil {
		fmt.Printf("Factory return error: %v\n", err)
		os.Exit(1)
	}

	r := p.Pay(200)

	fmt.Println(r)
}

//END5 OMIT
