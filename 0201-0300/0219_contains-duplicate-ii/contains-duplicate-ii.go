package main

/*
  Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

  Example 1:
    Input: nums = [1,2,3,1], k = 3
    Output: true

  Example 2:
    Input: nums = [1,0,1,1], k = 1
    Output: true

  Example 3:
    Input: nums = [1,2,3,1,2,3], k = 2
    Output: false

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. -10^9 <= nums[i] <= 10^9
    3. 0 <= k <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/contains-duplicate-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := map[int]int{}
	for i, num := range nums {
		if j, ok := dict[num]; ok && i-j <= k {
			return true
		}
		dict[num] = i
	}
	return false
}
