/*
假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。
由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。
给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
答案保证在 32 位有符号整数范围内
*/

package g502

import (
	"container/heap"
	"sort"
)

type Project struct {
	Profit  int
	Capital int
}

type ProjectSlice []Project

func (p ProjectSlice) Len() int           { return len(p) }
func (p ProjectSlice) Less(i, j int) bool { return p[i].Capital < p[j].Capital }
func (p ProjectSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type ProfitHeap struct {
	Data []int
}

func (p ProfitHeap) Len() int            { return len(p.Data) }
func (p ProfitHeap) Less(i, j int) bool  { return p.Data[i] < p.Data[j] }
func (p ProfitHeap) Swap(i, j int)       { p.Data[i], p.Data[j] = p.Data[j], p.Data[i] }
func (p *ProfitHeap) Push(v interface{}) { p.Data = append(p.Data, v.(int)) }
func (p *ProfitHeap) Pop() interface{} {
	data := p.Data
	v := data[len(data)-1]
	p.Data = data[:len(data)-1]
	return v
}

//IPO
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	var projects ProjectSlice = make(ProjectSlice, len(capital))

	for index, element := range profits {
		project := Project{
			Profit:  element,
			Capital: capital[index],
		}
		projects[index] = project
	}
	sort.Sort(projects)
	size := len(projects)
	var tempProfits = new(ProfitHeap)
	for start := 0; k > 0; k-- {

		for start < size && projects[start].Capital <= w {
			heap.Push(tempProfits, projects[start].Profit)
			start++
		}
		if tempProfits.Len() == 0 {
			break
		}

		w += heap.Pop(tempProfits).(int)
	}
	return w
}
