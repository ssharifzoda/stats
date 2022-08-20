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
			Status:   types.StatusFile,
		},
		{
			Amount:   13,
			Category: "Gorrila",
			Status:   types.StatusOK,
		},
		{
			Amount:   21,
			Category: "Esse change",
			Status:   types.StatusInProgress,
		},
	}
	fmt.Println(TotalInCategory(payments, "Esse change"))
	//Output: 21
}
