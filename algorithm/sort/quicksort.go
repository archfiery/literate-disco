package sort

import (
	"github.com/archfiery/literate-disco/container"
	"github.com/archfiery/literate-disco/error"
	"math/rand"
	"time"
)

func Quicksort(A []interface{}, f container.CompFunc) {
	quicksort(A, 0, len(A)-1, f)
}

func quicksort(A []interface{}, p, r int, f container.CompFunc) {
	if p < r {
		q, err := quicksortPartition(A, p, r, f)
		if err == nil {
			quicksort(A, p, q-1, f)
			quicksort(A, q+1, r, f)
		}
	}
}

func quicksortPartition(A []interface{}, p, r int, f container.CompFunc) (int, error.Error) {
	// use random
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(r-p) + p
	A[x], A[r] = A[r], A[x]
	// start partitioning
	i := p - 1
	for j := p; j < r; j++ {
		val, err := f(A[j], A[r])
		if err == nil && val == true {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]

	return i + 1, nil
}
