package umath

type numerical interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Max[T numerical](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i])
	}

	return res
}

func max[T numerical](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T numerical](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = min(res, nums[i])
	}

	return res
}

func min[T numerical](a, b T) T {
	if a < b {
		return a
	}

	return b
}
