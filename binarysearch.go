package main

func BinarySearch(nums []int, target int) int {
	var min int
	var max = len(nums) - 1
	var guess = 0

	for min <= max {
		guess = (min + max) / 2

		switch {
		case nums[guess] == target:
			return guess
		case nums[guess] < target:
			min = guess + 1
		default:
			max = guess - 1
		}
	}
	return -1
}
