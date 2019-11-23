package BinarySearch
// go test
import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array_slice := []int{5,3,7}
	for _, value := range  array_slice {
		result := BinarySearch(array_slice, value)
		if result == -1 {
			t.Fail()
		}
	}
}