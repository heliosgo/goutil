package umath

type numerical interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Max[T numerical](a, b T, others ...T) T {
	res := max(a, b)
	for _, v := range others {
		res = max(res, v)
	}

	return res
}

func max[T numerical](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T numerical](a, b T, others ...T) T {
	res := min(a, b)
	for _, v := range others {
		res = min(res, v)
	}

	return res
}

func min[T numerical](a, b T) T {
	if a < b {
		return a
	}

	return b
}
