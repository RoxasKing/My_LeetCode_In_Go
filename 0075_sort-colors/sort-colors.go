package main

/*
  给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
  此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
  注意:
  不能使用代码库中的排序函数来解决这道题。
*/

// 3-way partition
func sortColors(nums []int) {
	m, l, r := 0, 0, len(nums)-1
	for m <= r {
		if nums[m] == 0 {
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		} else if nums[m] == 1 {
			m++
		} else {
			nums[r], nums[m] = nums[m], nums[r]
			r--
		}
	}
}
