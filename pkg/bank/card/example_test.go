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
		types.Card {
			Balance: 10_000_00,
			Active: true,
		},

	}
	fmt.Println(Total(cards))
	//Output:3000000
	
}

func ExamplePaymentSources () {

	cards := []types.Card{
		types.Card {
			PAN: "5058 xxxx xxxx 8881",
			Balance: 10_000_00,
			Active: true,
		},
		types.Card {
			PAN: "5058 xxxx xxxx 8882",
			Balance: 20_000_00,
			Active: true,
		},
		types.Card {
			PAN: "5058 xxxx xxxx 8883",
			Balance: 30_000_00,
			Active: true,
		},

	}

	fmt.Println(PaymentSources(cards))
	//Output:
	//[{card 5058 xxxx xxxx 8881 1000000} {card 5058 xxxx xxxx 8882 2000000} {card 5058 xxxx xxxx 8883 3000000}]
}