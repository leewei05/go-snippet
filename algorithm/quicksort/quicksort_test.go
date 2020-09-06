package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testData := []int{4, 6, 3, 2, 5, 7, 9, 8, 1}
	expectedData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	quickSort(testData)

	ok := reflect.DeepEqual(expectedData, testData)
	if !ok {
		t.Errorf("fail of quickSort sort: [%v]", testData)
	}

	testData2 := []int{4, 6, 3, 2, 5, 7, 9, 8, 1}
	actual := quickSort2(testData2)

	ok = reflect.DeepEqual(expectedData, actual)
	if !ok {
		t.Errorf("fail of quickSort sort2: [%v]", actual)
	}

	testData3 := []int{4, 6, 3, 2, 5, 7, 9, 8, 1}
	quickSort3(testData3, 0, len(testData3)-1)

	ok = reflect.DeepEqual(expectedData, testData3)
	if !ok {
		t.Errorf("fail of quickSort sort2: [%v]", actual)
	}
}
