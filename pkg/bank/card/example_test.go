package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

//var daysInMonth = 30
//var daysInYear =365
//var percent = 3

func ExampleTotal() {
	cards := []types.Card{
		types.Card {
			Balance: 10_000_00,
			Active: true,
		},
		types.Card {
			Balance: 10_000_00,
			Active: true,
		},

	}
	fmt.Println(Total(cards))
	//Output:2000000
	
}