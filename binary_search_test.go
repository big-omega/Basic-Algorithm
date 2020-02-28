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

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   nil,
				target: -1,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0},
				target: -1,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0},
				target: 0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{0},
				target: 1,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: -1,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: 0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: 2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: 6,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: 1,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{0, 2, 5},
				target: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
