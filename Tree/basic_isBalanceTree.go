package Tree

import "math"

type RType struct {
	isBalance bool
	height    int
}

func basic_isBalanceTree(root *TreeNode) bool {
	res := pocess(root)
	return res.isBalance
}

func pocess(root *TreeNode) RType {
	if root == nil {
		return RType{
			isBalance: true,
			height:    0,
		}
	}
	// 记录左右的数据
	LeftData := pocess(root.Left)
	RightData := pocess(root.Right)

	// 记录
	// 高度为左右子树最大高度+1，是否平衡为左平衡右平衡且高度差的绝对值<2
	height := max(LeftData.height, RightData.height) + 1
	isBalance := LeftData.isBalance && RightData.isBalance && math.Abs(float64(LeftData.height-RightData.height)) < 2

	return RType{
		isBalance: isBalance,
		height:    height,
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
