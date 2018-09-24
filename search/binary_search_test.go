package search

import(
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 5, 6}
	result1 := BinarySearch( arr, 3)
	result2 := BinarySearch( arr, 2)

	if result1 == 2 && result2 == 1 {
		t.Log("Test ok")
	} else {
		t.Error("Test BinarySearch failt!")
	}
}