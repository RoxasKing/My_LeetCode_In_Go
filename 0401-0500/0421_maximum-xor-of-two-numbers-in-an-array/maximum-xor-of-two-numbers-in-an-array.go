package main

/*
  Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 ≤ i ≤ j < n.

  Follow up: Could you do this in O(n) runtime?

  Example 1:
    Input: nums = [3,10,5,25,2,8]
    Output: 28
    Explanation: The maximum result is 5 XOR 25 = 28.

  Example 2:
    Input: nums = [0]
    Output: 0

  Example 3:
    Input: nums = [2,4]
    Output: 6

  Example 4:
    Input: nums = [8,10,2]
    Output: 10

  Example 5:
    Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
    Output: 127

  Constraints:
    1. 1 <= nums.length <= 2 * 10^4
    2. 0 <= nums[i] <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
func findMaximumXOR(nums []int) int {
	out := 0
	t := NewTrie()
	for _, num := range nums {
		if res := t.Query(num); res > out {
			out = res
		}
		t.Insert(num)
	}
	return out
}

type Trie struct {
	child [2]*Trie
}

func NewTrie() *Trie {
	return &Trie{child: [2]*Trie{}}
}

func (t *Trie) Insert(num int) {
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if t.child[idx] == nil {
			t.child[idx] = NewTrie()
		}
		t = t.child[idx]
	}
}

func (t *Trie) Query(num int) int {
	out := 0
	for i := 31; i >= 0; i-- {
		if t == nil {
			return 0
		}
		idx := (num >> i) & 1
		if t.child[1-idx] != nil {
			out |= 1 << i
			t = t.child[1-idx]
		} else {
			t = t.child[idx]
		}
	}
	return out
}