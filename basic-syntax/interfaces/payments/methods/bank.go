package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBankMethod() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	fmt.Println("Производим оплату через банк, на сумму: ", usd)
	return rand.Int()
}

func (b Bank) Cancel(id int) {
	fmt.Println("Производим отмену операции через банк, id: ", id)
}
