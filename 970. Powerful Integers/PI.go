package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
}

func powerfulIntegers(x int, y int, bound int) []int {
	res := []int{}
	var xlist []int
	var ylist []int
	for i:=0;i<100;i++{
		x:=math.Pow(float64(x),float64(i))
		if int(x)>bound{
		break
		}
		xlist=append(xlist,int(x))
	}
	for i:=0;i<100;i++{
		y:=math.Pow(float64(y),float64(i))
		if int(y)>bound{
			break
		}
		ylist=append(ylist,int(y))
	}

	for _,xitem:=range xlist{
		for _,yitem:=range ylist{
			res=append(res,xitem+yitem)

		}
	}
	sort.Ints(res)
result:= map[int]bool{}
	for _,item :=range res{
		if item>bound {
			break
		}
		result[item]=true
	}
	res1:=[]int{}
	for key,_:=range result{
		res1=append(res1,key)
	}
	return res1
}
