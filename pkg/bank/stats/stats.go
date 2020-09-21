package stats

import "github.com/a1ishm/bank/pkg/bank/types"

// Avg f
func Avg(payments []types.Payment) types.Money {
	var i int
	var avg types.Money
	var sum types.Money
	for index, payment := range payments {
		sum += payment.Amount
		i = index
	}
	avg = sum / types.Money(i + 1)
	return avg
}	