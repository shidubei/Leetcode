package Arithmetic

import (
	"fmt"
	"sort"
)

// 贪心算法的基本思想题目
// 贪心算法思想：局部最优达到整体最优：即每次进行选择时选择每一次最优选择，从而实现整体效率最优
// 算法问题：局部最优不一定为整体最优解，但局部最优大概率是一个优良解

/*题目：给定一个项目列表，项目包含起始时间和结束时间，现在有一个会议室*/
/*会议室一次只能进行一个项目，要求找到一个列表，使得会议室一天能接纳的会议最多*/

type Program struct {
	Start int // 开始时间
	End   int // 结束时间
}
type Plist []Program

func (p Plist) Len() int {
	return len(p)
}
func (p Plist) Less(i, j int) bool {
	return p[i].End < p[j].End
}
func (p Plist) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TimeProgram(pList Plist, begin int) {
	// 将所有的program按照结束时间从早到晚排序
	sort.Sort(pList)
	res := []Program{}
	for _, p := range pList {
		if begin <= p.Start {
			res = append(res, p)
			begin = p.End
		}
	}
	fmt.Println(res)
}
