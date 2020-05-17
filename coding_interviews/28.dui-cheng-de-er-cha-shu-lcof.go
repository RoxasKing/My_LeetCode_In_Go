package codinginterviews

/*
  请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
  例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
      1
     / \
    2   2
   / \ / \
  3  4 4  3
  但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
      1
     / \
    2   2
     \   \
     3    3

  限制：
  0 <= 节点个数 <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func isEqual(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return isEqual(l.Left, r.Right) && isEqual(l.Right, r.Left)
}