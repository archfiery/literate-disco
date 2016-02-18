package sort

import (
	"fmt"
	"github.com/archfiery/literate-disco/container"
	"github.com/archfiery/literate-disco/error"
)

func Quicksort(A []interface{}, f container.CompFunc) {

}

func quicksortPartition(A []interface{}, p, r int, f container.CompFunc) (int, error.Error) {
	i := p - 1
	for j := p; j < r; j++ {
		val, err := f(A[j], A[r])
		if err == nil && val == true {
			i++
			A[i], A[j] = A[j], A[i]
		}
		fmt.Println(A)
	}
	A[i+1], A[r] = A[r], A[i+1]

	return i + 1, nil
}
