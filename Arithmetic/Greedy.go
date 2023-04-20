package Arithmetic

// 贪心算法的基本思想题目
// 贪心算法思想：局部最优达到整体最优：即每次进行选择时选择每一次最优选择，从而实现整体效率最优
// 算法问题：局部最优不一定为整体最优解，但局部最优大概率是一个优良解

/*题目：给定一个项目列表，项目包含起始时间和结束时间，现在有一个会议室*/
/*会议室一次只能进行一个项目，要求找到一个列表，使得会议室一天能接纳的会议最多*/

type Program struct {
	start int // 开始时间
	end   int // 结束时间
}
type PList struct {
	list []*Program //
}

func timeProgram(pList PList) {

}
