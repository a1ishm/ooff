package stats

import (
	"fmt"
	"github.com/a1ishm/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 10,
		},
		{
			ID:     2,
			Amount: 12,
		},
		{
			ID:     3,
			Amount: 16,
		},
		{
			ID:     4,
			Amount: 18,
		},
	}

	result := Avg(payments);
	fmt.Println(result)

	// Output: 14
}
