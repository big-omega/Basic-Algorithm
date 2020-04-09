package main_test

import "testing"

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func generateOrderedArray(size int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i
	}

	return nums
}

func Test_binarySearch(t *testing.T) {
	upBound := 1000000
	nums := generateOrderedArray(upBound)
	for i := range nums {
		if got := binarySearch(nums, i); got != i {
			t.Errorf("binarySearch() = %d, expected %d\n", got, i)
		}
	}

	outboundNums := []int{-5, -4, -3, -2, -1, upBound, upBound + 1, upBound + 2, upBound + 3}
	for _, v := range outboundNums {
		if got := binarySearch(nums, v); got != -1 {
			t.Errorf("binarySearch() = %d, expected %d\n", got, -1)
		}
	}
}
