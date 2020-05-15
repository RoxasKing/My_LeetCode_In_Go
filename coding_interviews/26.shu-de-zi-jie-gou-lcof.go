package codinginterviews

/*
  输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
  B是A的子结构， 即 A中有出现和B相同的结构和节点值。

  例如:
  给定的树 A:
       3
      / \
     4   5
    / \
   1   2
  给定的树 B：
     4
    /
   1
  返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	queue := []*TreeNode{A}
	for len(queue) != 0 {
		A, queue = queue[0], queue[1:]
		if isSubTree(A, B) {
			return true
		}
		if A.Left != nil {
			queue = append(queue, A.Left)
		}
		if A.Right != nil {
			queue = append(queue, A.Right)
		}
	}
	return false
}

func isSubTree(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return isSubTree(a.Left, b.Left) && isSubTree(a.Right, b.Right)
}
