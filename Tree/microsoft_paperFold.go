package Tree

import "fmt"

/*题目：给一张长纸条，每次对折产生一个折痕，如对折一次产生一个凹折痕（纸条每次上折），问对折n次，所有折痕的方向是什么*/

/*
思路：
将长纸条看作一个二叉树，二叉树的左孩子永远是凹，右孩子永远是凸
*/
func Microsoft_paperFolding(i, n int, isLeft bool) {
	if i > n {
		return
	}
	Microsoft_paperFolding(i+1, n, true)
	if isLeft == true {
		fmt.Print("凹 ")
	} else {
		fmt.Print("凸 ")
	}
	Microsoft_paperFolding(i+1, n, false)
}
