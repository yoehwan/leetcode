package main

import (
	"fmt"
)

func main() {
fmt.Println(minFlips("001011101"))
}

func minFlips(target string) int {

	//Check foreword number is zero
	for target[0]==48{
		if len(target)==1{
			return 0
		}
		target=target[1:]
	}
	if target=="1"{
		return 1
	}
	//make boolList
	fmt.Println(target)
	res:=[]byte{49}
	for i:=1;i<len(target);i++{
		if target[i]!=res[len(res)-1]{
			if target[i]==49 {
				res = append(res, 49)
				continue
			}
				res=append(res,48)
			}
		}

	return len(res)
	}

