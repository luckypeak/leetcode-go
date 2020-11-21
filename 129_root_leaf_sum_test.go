package leetcode_go

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		"test",
		args{root: root},
		12,
	},
		{
			"test",
			args{root: root2},
			25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
