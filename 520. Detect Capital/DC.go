package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(detectCapitalUse("usA"))
}

func detectCapitalUse(word string) bool {
	 if len(word)==1{
		return true
	}
//A=65,a=97,Z=90,z=122
	list:=[]bool{}
	if word[0]<90{
	for i:=1;i<len(word);i++{
			if word[i]<=90{
				list=append(list,true)
			}else {
				list=append(list,false)
			}
		}
	return checkAll(list)
	}
return strings.ToLower(word)==word
}

func checkAll(data []bool) bool {
	for _,item:=range data{
		if data[0]!=item{
			return false
		}
	}
	return true
}

/*
func detectCapitalUse(word string) bool {
    return strings.ToLower(word) == word || strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}
*/