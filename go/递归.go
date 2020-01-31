/*
	1. 终止条件
	2. 本级递归需要做什么
	3. 返回值是什么，给上一级返回什么信息
	递归即程序反复调用自身，反复调用自身，每级的处理相同，只需关注一级递归的解决过程即可

*/

package main

/*
	反转单链表 1->2->3->4 得到  4->3->2->1

*/
type Node struct {
	data int
	next *Node
}

func reverseList(head Node) Node {
	if head == nil || head.next == nil {
		return head
	}
	/*
		234，reverseList后得到432，1仍指向2
		此时将2指向1，1放到末尾即可
	*/
	newList := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return head
}

/*
	二叉树的最大深度
	https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	/*
		退出条件： 没有叶子节点，即 left==nil && right==nil
		本级返回： 本级深度
		本级处理： 本机深度=下一级+1
		递归后，树变成root、root.left、root.right，其中root.left和root.right分别记录的是root的左右子树的最大深度。
	*/

	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return int(math.Max(float64(left), float64(right))) + 1
}

/*
	两辆交换链表中节点
	https://leetcode-cn.com/problems/swap-nodes-in-pairs/
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	/*
		结束条件： 链表为空或链表下一个为空
		本次处理： head head.Next newNextHead(已处理好的部分)
				  head.Next.Next=head
				  head.Next=newNextHead
				  上一级->head->next->已经处理好的下一级
				  上一级->next->head->已经处理好的下一级
		返回：    处理好的head
	*/
	if head == nil || head.Next == nil {
		return head
	}
	nextList := swapPairs(head.Next.Next)
	head.Next.Next = head
	head.Next = nextList
	return head
}

/*
	平衡二叉树：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1
	https://leetcode-cn.com/problems/balanced-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
	终止条件： 树为空
	本级执行： 判断本级，条件： 左右子树均为平衡二叉树，且高度相差不大于一
	返回： 本级是否是，左子树、右子树最大深度
*/
func isBalanced(root *TreeNode) bool {
	_, b := isBWithHeigth(root)
	return b
}

func isBWithHeigth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lH, lB := isBWithHeigth(root.Left)
	rH, rB := isBWithHeigth(root.Right)
	isB := lB && rB && math.Abs(float64(lH)-float64(rH)) <= 1
	return int(math.Max(float64(lH), float64(rH))) + 1, isB
}

/*
	二叉树的最小深度
	https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	/*注意left == nil && root == nil才是叶子节点*/
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	return int(math.Min(float64(l), float64(r))) + 1
}

/*
	翻转二叉树
	https://leetcode-cn.com/problems/invert-binary-tree/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	/*
	   退出条件：root == nil
	   本次处理：root left right，left和right代表已经换完的
	       返回：root
	*/
	if root == nil {
		return root
	}
	l := invertTree(root.Left)
	r := invertTree(root.Right)
	root.Left, root.Right = r, l
	return root
}

/*
	合并二叉树：
	https://leetcode-cn.com/problems/merge-two-binary-trees/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	/*
	   退出：t1 ==nil || t2 == nil，return t1 t2中非nil的，都为 nil return nil
	   处理：t1 t2 newl,newr,new=t1+t2
	   返回：new
	*/
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	newl := mergeTrees(t1.Left, t2.Left)
	newr := mergeTrees(t1.Right, t2.Right)
	new := &TreeNode{
		Left:  newl,
		Right: newr,
		Val:   t1.Val + t2.Val,
	}

	return new
}

/*
	最大二叉树
	https://leetcode-cn.com/problems/maximum-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	/*
			返回： len(nums)==0 return nil
				  len(nums)==1 return &TreeNode{Val:nums[0]}
		    处理： 找到最大的，左树，右树
		    返回： 处理好的node
	*/
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	maxKey := 0
	maxVal := nums[0]
	for k, v := range nums {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}
	l := constructMaximumBinaryTree(nums[:maxKey])
	r := constructMaximumBinaryTree(nums[maxKey+1:])
	return &TreeNode{Val: maxVal, Left: l, Right: r}
}

/*
	删除链表中的重复元素
	https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	/*
		结束： head==nil || head.Next ==nil, return head
		处理： head next，next是处理好的，head.Val 和 next.Val 相等，则head指向next的next
		返回： head
	*/
	if head == nil || head.Next == nil {
		return head
	}
	newNode := deleteDuplicates(head.Next)
	if head.Val == newNode.Val {
		head.Next = newNode.Next
	}
	return head
}

/*
	countandsay
	https://leetcode-cn.com/problems/count-and-say/
*/

func countAndSay(n int) string {
	/*
		结束： n=0||n=1
		处理： 得到后面的string，找连续数字，拼接
		返回： 本级字符串
	*/
	if n == 0 {
		return ""
	} else if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	}
	str := countAndSay(n - 1)
	i, k := 1, 0
	count := 1
	result := ""
	for i = 1; i < len(str); i++ {
		if str[k] == str[i] {
			count++
		} else {
			result = fmt.Sprintf("%s%d%s", result, count, string(str[k]))
			count = 1
		}
		k++
	}
	result = fmt.Sprintf("%s%d%s", result, count, string(str[k]))

	return result
}

/*
	全排列
	https://leetcode.com/problems/permutations/
	[1,2,3]
	找到后面子列的全排列，当前元素插入道每个间隔
*/

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	before := permute(nums[:len(nums)-1])
	new := make([][]int, 0, len(before))
	cur := nums[len(nums)-1]

	for _, v := range before {
		fmt.Println(v)
		for k := 0; k < len(v); k++ {
			one := make([]int, 0, len(v)+1)
			tmp := make([]int, len(v[0:k]))
			copy(tmp, v[0:k])
			one = append(one, tmp...)
			one = append(one, cur)
			one = append(one, v[k:]...)
			new = append(new, one)
		}
		new = append(new, append(v, cur))
	}
	return new
}

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	sublist := permute(nums[1:])
	cur := nums[0]
	result := make([][]int, 0, len(100))
	for _, list := range sublist {
		for k, _ := range list {
			tmp := make([]int, 0, len(sublist)+1)
			tmp = append(tmp, list[:k])
			tmp = append(tmp, cur)
			tmp = append(tmp, list[k:])
			result = append(result, tmp)
		}
	}
	return result
}

/*
	回溯算法
*/
func permute1(nums []int) [][]int {

}

func traceback(nums, temp []int, key int) {

	for i := 0; i < len(nums); i++ {
		traceback(nums, temp, key)
	}
}

/*
	Combination Sum
	https://leetcode.com/problems/combination-sum/
*/

var res [][]int
var res1 []int

func combinationSum(candidates []int, target int) [][]int {
	submission(candidates, target, len(candidates)-1, res1)
	return res
}

func submission(candidates []int, target int, index int, res1 []int) {
	if target == 0 {
		res = append(res, res1)
		return
	}
	for index >= 0 {
		if candidates[index] <= target {
			res = append(res, candidates[index])
			submission(candidates, target-candidates[index], index, res1)
			res1 = res[:len(res1)-1]
		}
		index--
	}
	return
}
