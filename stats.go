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
func CategoryAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	var category types.Category
	var sum, count types.Money
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
			count++
			categories[category] = sum / count
		}
		if payment.Category != category {
			category = payment.Category
			count = 1
			sum = payment.Amount
			categories[category] = sum / count
		}
	}
	return categories
}
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	results := map[types.Category]types.Money{}
	for category, amount := range second {
		results[category] = amount
	}
	var scht int
	var indikator types.Category
	for category, amount := range first {
		scht++
		if scht <= len(results) {
			for category2, amount2 := range results {
				if category2 == category {
					results[category2] = amount2 - amount
					indikator = category2
					break
				}
				continue
			}
			if indikator == category {
				continue
			}
			results[category] = 0 - amount
		}
	}
	return results
}
