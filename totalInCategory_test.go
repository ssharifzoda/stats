package stats

import (
	"fmt"
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   20,
			Category: "Esse change",
		},
		{
			Amount:   13,
			Category: "Gorrila",
		},
		{
			Amount:   21,
			Category: "Esse change",
		},
	}
	fmt.Println(TotalInCategory(payments, "Esse change"))
	//Output: 41
}
