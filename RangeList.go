package main

import (
	"fmt"
	_ "fmt"
)

type RangeList struct {
	 list [][]int
}

func (rangeList *RangeList) Add(rangeElement [2]int) error {
	rangeList.list=insert(rangeList.list,rangeElement)
	return nil
}

func (rangeList *RangeList) Remove(rangeElement [2]int) error {
	rangeList.list=delete(rangeList.list,rangeElement)
	return nil
}

func (rangeList *RangeList) Print() error {

	for _, lists := range rangeList.list{

		str := lists[0]
		str1 :=lists[1]
		fmt.Print("  [ ")
		fmt.Print(str)
		fmt.Print(" , ")
		fmt.Print(str1)
		fmt.Print(" )")
	}
	fmt.Println()
	 return nil
}

func insert(intervals [][]int, newInterval [2]int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {		// 在插入区间的右侧且无交集

			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {			// 在插入区间的左侧且无交集

			ans = append(ans, interval)			// 与插入区间有交集，计算它们的并集
		} else {

			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

func delete(intervals [][]int, newInterval [2]int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	if(left<right){
		for _, interval := range intervals {
			if interval[0] > right || interval[1] < left{
				ans = append(ans, interval)                         // 与删除区间无交集,则直接添加数值对
			} else if !(interval[0]>=left && interval[1]<right){	//若不被删除区间完全包围
				if interval[0]<left && interval[1]>left {
					ans=append(ans,[]int{interval[0],left})         //添加左侧的被部分删除的新区间
				}
				if interval[0]<right && interval[1]>right {
					ans=append(ans,[]int{right,interval[1]})        //添加右侧的被部分删除的新区间
				}
			}
		}
	}else{
		ans=intervals
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {

	rl := RangeList{}

	rl.Add([2]int{1, 5})
	rl.Print()

	rl.Add([2]int{10, 20})
	rl.Print()

	rl.Add([2]int{20, 20})
	rl.Print()

	rl.Add([2]int{20, 21})
	rl.Print()


	rl.Add([2]int{2, 4})
	rl.Print()

	rl.Add([2]int{3, 8})
	rl.Print()

	rl.Remove([2]int{10, 10})
	rl.Print()

	rl.Remove([2]int{10, 11})
	rl.Print()

	rl.Remove([2]int{15, 17})
	rl.Print()

	rl.Remove([2]int{3, 19})
	rl.Print()
}
