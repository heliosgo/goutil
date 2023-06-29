package uslice

// Determine whether an element exists
func Exists[T comparable](sli []T, x T) bool {
	for _, v := range sli {
		if v == x {
			return true
		}
	}

	return false
}

func Existsf[T any](sli []T, f func(x T) bool) bool {
	for _, v := range sli {
		if f(v) {
			return true
		}
	}

	return false
}
