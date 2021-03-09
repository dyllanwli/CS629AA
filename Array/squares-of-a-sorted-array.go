package main

func sortedSquares(nums []int) []int {
	// two points start from the head and tail of the array at the same time
	// one the two points, it will end
	res := make([]int, len(nums))
	l := len(nums) - 1
	for h, k, t := 0, l, l; h <= t; k-- {
		if nums[h]*nums[h] > nums[t]*nums[t] {
			res[k] = nums[h] * nums[h]
			h++
		} else {
			res[k] = nums[t] * nums[t]
			t--
		}
	}
	return res
}
