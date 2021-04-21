package main

/*
  Given an integer array nums and an integer k, return true if nums has a continuous subarray of size at least two whose elements sum up to a multiple of k, or false otherwise.

  An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.

  Example 1:
    Input: nums = [23,2,4,6,7], k = 6
    Output: true
    Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

  Example 2:
    Input: nums = [23,2,6,4,7], k = 6
    Output: true
    Explanation: [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
    42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.

  Example 3:
    Input: nums = [23,2,6,4,7], k = 13
    Output: false

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 0 <= nums[i] <= 10^9
    3. 0 <= sum(nums[i]) <= 2^31 - 1
    4. 1 <= k <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/continuous-subarray-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum + Hash
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	dict := map[int]int{0: -1}
	sum := 0
	for r, num := range nums {
		sum += num
		sum %= k
		if l, ok := dict[sum]; ok {
			if r-l >= 2 {
				return true
			}
		} else {
			dict[sum] = r
		}
	}
	return false
}
