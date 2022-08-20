package stats

import (
	"fmt"
	types "github.com/ssharifzoda/kun/v2"
)

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   20,
			Category: "Esse change",
			Status:   "FAIL",
		},
		{
			Amount:   13,
			Category: "Gorrila",
		},
		{
			Amount:   21,
			Category: "Esse change",
			Status:   "OK",
		},
	}
	fmt.Println(TotalInCategory(payments, "Esse change"))
	//Output: 21
}
