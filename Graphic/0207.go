package Graphic

import "fmt"

/*题目：给一个课程表二维数组，每一个数组的[ai,bi]表示要完成bi，需要先完成ai,问是狗能完成所有的课*/
/*思路：转化为图，拓补排序*/

func CanFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	// 记录所有出现过的点,看课程是否存在
	isExist := make(map[int]bool)
	for i := 0; i < len(prerequisites); i++ {
		for j := 0; j < len(prerequisites[0]); j++ {
			isExist[prerequisites[i][j]] = true
		}
	}
	fmt.Println(isExist)
	// inMap用于记录所有节点的入度数
	// 节点即数组中位于第二列的点，每出现一次，入度加1
	inMap := make(map[int]int)
	// 对于存在的点，先将他们的入度都置为0
	for i, v := range isExist {
		if v {
			inMap[i] = 0
		}
	}
	// 如果这些点在[i][0]位出现，有入度，入度++
	for i := 0; i < len(prerequisites); i++ {
		inMap[prerequisites[i][0]]++
	}
	fmt.Println(inMap)
	// 0入度队列，所有0入度的都优先入队，先进先出
	zeroInQueue := []int{}
	for i, v := range inMap {
		// 查找入度为0的点
		if v == 0 {
			zeroInQueue = append(zeroInQueue, i)
		}
	}
	fmt.Println(zeroInQueue)
	if len(zeroInQueue) == 0 {
		return false
	}
	// 当zeroInQueue中存在数据的时候，说明有入度为0的节点
	for len(zeroInQueue) > 0 {
		// 队列弹出
		cur := zeroInQueue[0]
		zeroInQueue = zeroInQueue[1:]
		inMap[cur] = -1
		// inMap里以cur为入度的点入度-1
		// 查找以cur为入度的点，即[i][0]==cur
		for i := 0; i < len(prerequisites); i++ {
			if prerequisites[i][0] == cur {
				inMap[prerequisites[i][1]]--
			}
			fmt.Println(inMap)

		}
		for i, v := range inMap {
			if v == 0 {
				zeroInQueue = append(zeroInQueue, i)
			}
		}
	}
	for _, v := range inMap {
		if v != -1 {
			return false
		}
	}
	return true
}

//func CanFinish2(numCourses int, prerequisites [][]int) bool {
//	// 邻接表
//	afterMap := map[int][]int{}
//	inMap := make(map[int]int,numCourses)
//
//	for _,p := range prerequisites {
//		after,first := p[]
//	}
//}
