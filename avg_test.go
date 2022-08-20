package stats

import (
	"fmt"
	types "github.com/ssharifzoda/kun/v2"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 1000,
			Status: "OK",
		},
		{
			Amount: 200,
			Status: "FAIL",
		},
	}
	fmt.Println(Avg(payments))
	//Output:500
}
