package main

import (
	"fmt"	
	"sort"
)


func main() {
	test :=[][]int{{3,0,8,4},{2,4,5,7},{9,2,6,3},{0,3,1,0}}
	fmt.Print(maxIncreaseKeepingSkyline(test))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	res:=sum(grid)
	row,col:=getRowCol(grid)




	for condition:=0;condition<len(grid);condition++{
			for check:=0;check<len(grid);check++{
				if col[condition]>=row[check]{
					grid[check][condition]=row[check]
				}else{
					grid[check][condition]=col[condition]
				}
			}
	}
        fmt.Print(grid)
	return sum(grid)-res
}

func sum(input [][]int)int{
 var res int
	for _,item:=range input{
		for _,item:=range item{
			res+=item
		}
	}
	
	return res
}

func getRowCol(input [][]int) ([]int,[]int) {
	var row,col []int
	for i:=0;i<len(input);i++{
		var _temp []int
		for j:=0;j<len(input);j++{
		_temp=append(_temp,input[j][i])
		}
		sort.Ints(_temp)
		col=append(col,_temp[len(input)-1])
	}
		
	for _,item:=range input{
	sort.Ints(item)	
	row=append(row,item[3])
	}
		return row,col
}