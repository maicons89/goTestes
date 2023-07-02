package ordering

import (
	"sort"
)

func OrderArray(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
