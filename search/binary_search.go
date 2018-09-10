package search

import(
	"fmt"
	"time"
)

func searchBinary(arr *[]int, l, r, gt int) int {
	if l == r {
		if l == 0 && gt < (*arr)[0] {
			return -1
		}
		return l
	}
	mid := (l+r)/2
	if gt > (*arr)[mid] {
		return searchBinary(arr, mid+1, r, gt)
	}
	return searchBinary(arr, l, mid, gt)
}

func searchBinaryB(arr *[]byte, l, r int, gt byte) int {
	if l == r {
		if l == 0 && gt < (*arr)[0] {
			return -1
		}
		return l
	}
	if l == r-1 {
		if (*arr)[l] > gt {
			return l-1
		}
		if (*arr)[r] > gt {
			return l
		}
		return r
	}
	mid := (l+r)/2
	// fmt.Println("\n", string(gt), l, r, string((*arr)[mid]))
	if gt <= (*arr)[mid] {
		return searchBinaryB(arr, l, mid, gt)
	}
	return searchBinaryB(arr, mid, r, gt)
}

type searchArray interface {
	BinarySearch() int
}

type intSearch struct {
	gt  int
	arr []int
}

func (this intSearch) BinarySearch() int {
	return -1
}

func BinarySearch (arr []int, gt int) int {
	if len(arr) == 0 {
		return -1
	}
	return searchBinary(&arr, 0, len(arr)-1, gt)
}

func BinarySearchB (arr []byte, gt byte) int {
	if len(arr) == 0 {
		return -1
	}
	return searchBinaryB(&arr, 0, len(arr)-1, gt)
}

func main2() {
	startTime := time.Now()
	a := []int{}
	for i := 2; i < 50000000; i++ {
		a = append(a, i)
	}
	for i := 0; i < 1000000; i++ {
		searchBinary(&a, 0, len(a)-1, 4)
	}


	// for i, gt := range a {
		fmt.Println(0, 0, searchBinary(&a, 0, len(a), 0))
	// }

	fmt.Println("~~~ END ~~~", time.Since(startTime))
}

func main() {
	for i:= 0; i < 2; i++ {
		fmt.Println("Test ", i+1)
		main2()
	}
}
