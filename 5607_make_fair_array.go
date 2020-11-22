package leetcode_go

func waysToMakeFair(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ans int
	oddNums := make([]int, len(nums))
	ennNums := make([]int, len(nums))
	oddNums[0] = nums[0]
	ennNums[0] = 0
	for i := 1; i < len(nums); i++ {
		if i%2 == 0 {
			oddNums[i] = oddNums[i-1] + nums[i]
			ennNums[i] = ennNums[i-1]
		} else {
			oddNums[i] = oddNums[i-1]
			ennNums[i] = ennNums[i-1] + nums[i]
		}
	}
	length := len(nums) - 1
	for index, num := range nums {
		if index%2 == 0 {
			if oddNums[index]+ennNums[length]-ennNums[index]-num == ennNums[index]+oddNums[length]-oddNums[index] {
				ans++
			}
		} else {
			if oddNums[index]+ennNums[length]-ennNums[index]+num == ennNums[index]+oddNums[length]-oddNums[index] {
				ans++
			}
		}
	}
	return ans
}

func isFair(nums []int) bool {
	var odd, enn int
	for index, num := range nums {
		if index%2 == 0 {
			odd += num
		} else {
			enn += num
		}
	}
	return odd == enn
}
