package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Print(isHappy(7))

}

func isHappy(n int)bool{
	testTime:=0
	
	if n==1{
		return true
	}

	for testTime<1000{
		if n==1{
			return true
		}
		n=loopNumber(n)
		testTime++
	}

	return false
}

func loopNumber(num int) int {
	var res int
	numstr:=strconv.Itoa(num)
		for _,item:=range numstr{
			tem,_:=strconv.Atoi(string(item))
			res+=tem*tem
		}
	
	return res
}