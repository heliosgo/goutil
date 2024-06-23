package umath

import "goutil/utype"

func Max[T utype.Number](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i])
	}

	return res
}

func max[T utype.Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T utype.Number](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = min(res, nums[i])
	}

	return res
}

func min[T utype.Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}
