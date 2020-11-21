package leetcode_go


var ans = 0
/**
  使用dfs 遍历根到每个叶子节点的路径，并记录过程中的路径。到达叶子节点加和
 */
func sumNumbers(root *TreeNode) int {
	ans = 0
	if root == nil{
		return 0
	}

	return bfs(root)
}

// 使用bfs，两个队列，一个队列放节点，一个队列放数据
func bfs(node *TreeNode) int {
	queue := make([]*TreeNode, 0)
	numQueue := make([]int , 0)
	queue = append(queue, node)
	sum := 0
	num := 0
	numQueue = append(numQueue, node.Val)
	for len(queue) > 0{
		node = queue[0]
		queue = queue[1:]
		num = numQueue[0]
		numQueue = numQueue[1:]
		if node.Left == nil && node.Right == nil{
			sum += num
		}else{
			if node.Left != nil{
				queue = append(queue, node.Left)
				numQueue = append(numQueue, num * 10 + node.Left.Val)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
				numQueue = append(numQueue, num * 10 + node.Right.Val)
			}
		}
	}

	return sum
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