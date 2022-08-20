package stats

func Avg(payments []types.Payment) types.Money {
	var sum int
	for _, payment := range payments {
		sum += int(payment.Amount)
	}
	return types.Money(sum / len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sumPayments int
	for _, payment := range payments {
		if category == payment.Category {
			sumPayments += int(payment.Amount)
		}
	}
	return types.Money(sumPayments)
}
