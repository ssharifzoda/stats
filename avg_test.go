package stats

import (
	"fmt"
	types "github.com/ssharifzoda/kun/v2"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 1000,
			Status: types.StatusOK,
		},
		{
			Amount: 2000,
			Status: types.StatusInProgress,
		},
		{
			Amount: 200,
			Status: types.StatusFile,
		},
	}
	fmt.Println(Avg(payments))
	//Output:1500
}
