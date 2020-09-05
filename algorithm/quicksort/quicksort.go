package quicksort

import (
	"math/rand"
)

// Time O(nlogn)
func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
}

func quickSort2(data []int) []int {
	if len(data) < 2 {
		return data
	}

	pivot := data[0]
	smaller := make([]int, 0, len(data))
	equal := make([]int, 1, len(data))
	equal[0] = pivot
	larger := make([]int, 0, len(data))
	for i := 1; i < len(data); i++ {
		if data[i] > pivot {
			larger = append(larger, data[i])
		} else if data[i] < pivot {
			smaller = append(smaller, data[i])
		} else {
			equal = append(equal, data[i])
		}
	}
	return append(append(quickSort2(smaller), equal...), quickSort2(larger)...)
}
