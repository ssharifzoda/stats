package stats

import types "github.com/ssharifzoda/kun/v2"

func Avg(payments []types.Payment) types.Money {
	var sum int
	var count int
	for _, payment := range payments {
		if payment.Status != types.StatusFile {
			sum += int(payment.Amount)
			count++
		}
	}
	return types.Money(sum / count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sumPayments int
	for _, payment := range payments {
		if payment.Status != types.StatusFile {
			if category == payment.Category {
				sumPayments += int(payment.Amount)
			}
		}
	}
	return types.Money(sumPayments)
}
