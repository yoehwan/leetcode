package main

import (
	"fmt"
)

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
}

func powerfulIntegers(x int, y int, bound int) []int {
	 trueMap:= map[int]bool{}
	xlist := powList(x, bound)
	ylist := powList(y, bound)
	for _, xitem := range xlist {
		for _, yitem := range ylist {
			if xitem+yitem <= bound {
				trueMap[xitem+yitem] = true
			}
		}
	}
		res:=[]int{}
	for key,value:=range trueMap{
		if value{
			res=append(res,key)
		}
	}
	return res
}

func powList(target int, bound int) []int {
	list := []int{1}
	for i := 1; i < 100; i++ {
		_temp := intPow(target, i)
		if _temp > bound {
			break
		}
		list = append(list, _temp)
	}

	return list
}
func intPow(target int, times int) int {
	res := target
	for i := 1; i < times; i++ {
		res = res * target
	}
	return res

}
