package main

/*
  给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

  说明:
  你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
*/

func missingNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i && nums[i] < len(nums) {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}