package leetcode_go


var ans = 0
func sumNumbers(root *TreeNode) int {
	ans = 0
	if root == nil{
		return 0
	}
	dfs(root, 0)
	return ans
}


func dfs(node *TreeNode, sum int) int {
	if node == nil{
		return sum
	}
	sum = sum *10  +  node.Val
	if node.Left == nil && node.Right == nil{
		ans += sum
		return sum
	}else {
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
}