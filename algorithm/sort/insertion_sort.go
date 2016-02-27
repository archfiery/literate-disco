package sort

import (
	"github.com/archfiery/literate-disco/common"
)

func BinaryInsertionSort(A []interface{}, n int, f common.CompFunc) {
	for i := 1; i < n; i++ {
		lo, hi := 0, i
		m := i / 2

		if v, err := f(A[i], A[m]); err == nil && v == false {
			lo = m + 1
		} else if v, err := f(A[i], A[m]); err == nil && v == false {
			hi = m
		}
		m = lo + (hi-lo)/2

		for lo < hi {
			if v, err := f(A[i], A[m]); err == nil && v == false {
				lo = m + 1
			} else if v, err := f(A[i], A[m]); err == nil && v == true {
				hi = m
			} else {
				break
			}

			m = lo + (hi-lo)/2
		}

		if m < i {
			tmp := A[i]
			for j := i - 1; j >= m; j-- {
				A[j+1] = A[j]
			}
			A[m] = tmp
		}
	}
}

func binarySearch(A []interface{}, lo int, hi int, key interface{}, f common.CompFunc) int {
	if lo == hi {
		return lo
	}
	var m int
	m = lo + (hi-lo)/2
	v, err := f(lo, hi)
	if err == nil && v == true {
		return binarySearch(A, m+1, hi, key, f)
	} else if err == nil && v == false {
		return binarySearch(A, lo, m, key, f)
	}

	return m
}
