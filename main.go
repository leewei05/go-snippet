package main

func main() {}

// Reverse slice of integers
func reverseSliceInt(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Check empty slice
func checkEmptySlice(eSlice []byte) {
	if len(eSlice) == 0 {
		// Do something if the slice is empty
	}
}

// Reverse string
func reverse(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
