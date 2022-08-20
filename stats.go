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
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtred []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtred = append(filtred, payment)
		}
	}
	return filtred
}
