package bubble

import "fmt"

func bubbleSort(slice []int) []int {
	fmt.Println("Bubble Sort")
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
