package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1,1,4,2,1,3}))
}

func heightChecker(height []int) int {
	res:=0
	arr:=[]int{}
	arr=append(arr,height...)
	sort.Ints(height)
	for idx,_:=range height{
		if arr[idx]!=height[idx]{
			res++
		}
	}
	return res
}
