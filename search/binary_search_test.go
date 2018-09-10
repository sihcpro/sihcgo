package search

import(
	"testing"
	"errors"
)

func TestBinarySearch(t *testing.T) {
	result, text := BinarySearch(false, dataPath)

	if !result {
		t.Error(errors.New(text))
	} else {
		t.Log("Tested "+text+" files")
	}
}