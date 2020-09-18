package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
	
		sum += card.Balance 
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	// TODO: ваш код
	var source []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		source = append(source, types.PaymentSource {
				Type: "card",
				Number: string(card.PAN),
				Balance: card.Balance,})
				
	}
	for _, sour := range source {
		fmt.Println(sour.Number)
	}
	return source
 }