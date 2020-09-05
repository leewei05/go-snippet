package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testData := []int{4, 6, 3, 2, 5, 7, 9, 8, 1}
	expectedData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	bubbleSort(testData)

	ok := reflect.DeepEqual(expectedData, testData)
	if !ok {
		t.Errorf("fail of bubble sort: [%v]", testData)
	}
}
