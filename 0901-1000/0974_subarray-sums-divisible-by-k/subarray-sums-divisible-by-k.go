package main

/*
  Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.

  Example 1:
    Input: A = [4,5,0,-2,-3,1], K = 5
    Output: 7
    Explanation: There are 7 subarrays with a sum divisible by K = 5:
    [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

  Note:
    1. 1 <= A.length <= 30000
    2. -10000 <= A[i] <= 10000
    3. 2 <= K <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Hash + Congruence Theorem
func subarraysDivByK(A []int, K int) int {
	out := 0
	dict := make(map[int]int)
	dict[0] = 1
	sum := 0
	for _, a := range A {
		sum += a
		sum %= K
		if sum < 0 {
			sum += K
		}
		out += dict[sum]
		dict[sum]++
	}
	return out
}