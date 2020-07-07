package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	
	monthTable:=[]int{0,31,28,31,30,31,30,31,31,30,31,30,31}
	splitData := strings.Split(date, "-")
	year,_:=strconv.Atoi(splitData[0])
	month,_:=strconv.Atoi(splitData[1])
	day,_:=strconv.Atoi(splitData[2])

	sum:=0
	
	if month == 1 {
		return day
	} else {
			
		if (year%4 == 0) && (year%100 != 0) || (year%400 == 0) {
			monthTable[2]=29
		}
		for i:=1;i<month;i++{
				sum+=monthTable[i]
			}
	}
	
	sum+=day
	return sum
}

func main() {
	fmt.Println(dayOfYear("2004-03-01"))

}
