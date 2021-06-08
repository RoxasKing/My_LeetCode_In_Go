package main

import "sort"

/*
  You have a cubic storeroom where the width, length, and height of the room are all equal to n units. You are asked to place n boxes in this room where each box is a cube of unit side length. There are however some rules to placing the boxes:

    1. You can place the boxes anywhere on the floor.
    2. If box x is placed on top of the box y, then each side of the four vertical sides of the box y must either be adjacent to another box or to a wall.

  Given an integer n, return the minimum possible number of boxes touching the floor.

  Example 1:
    Input: n = 3
    Output: 3
    Explanation: The figure above is for the placement of the three boxes.
      These boxes are placed in the corner of the room, where the corner is on the left side.

  Example 2:
    Input: n = 4
    Output: 3
    Explanation: The figure above is for the placement of the four boxes.
      These boxes are placed in the corner of the room, where the corner is on the left side.

  Example 3:
    Input: n = 10
    Output: 6
    Explanation: The figure above is for the placement of the ten boxes.
      These boxes are placed in the corner of the room, where the corner is on the back side.

  Constraints:
    1 <= n <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/building-boxes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func minimumBoxes(n int) int {
	l, r := 1, n
	for l < r {
		floorBoxes := l + (r-l)>>1
		if maximumBoxes(floorBoxes) < n {
			l = floorBoxes + 1
		} else {
			r = floorBoxes
		}
	}
	return l
}

func maximumBoxes(floorBoxes int) int {
	out := 0
	for floorBoxes > 0 {
		out += floorBoxes
		i := sort.Search(floorBoxes, func(i int) bool { return i*(i+1)>>1 > floorBoxes }) - 1
		next := i * (i - 1) >> 1
		if i*(i+1)>>1 < floorBoxes {
			next += floorBoxes - i*(i+1)>>1 - 1
		}
		floorBoxes = next
	}
	return out
}