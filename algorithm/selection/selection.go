package selection

import "fmt"

func selectSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[min] > slice[j] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
		fmt.Println(slice)
	}

	return slice
}
