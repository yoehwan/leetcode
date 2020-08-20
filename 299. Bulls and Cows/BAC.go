package main

import (
	"fmt"
	"strconv"
)

func main() {
	secret := "1807"
	guess := "7810"
	fmt.Println(getHint(secret, guess))

}

func getHint(secret string, guess string) string {
	A := 0
	B := 0
	//length:= len(secret)
	slist := map[uint8]int{}
	glist := map[uint8]int{}
	for idx, _ := range secret {
		if secret[idx] == guess[idx] {
			A++
			continue
		}
		slist[secret[idx]]++
		glist[guess[idx]]++
	}
	for key, _ := range slist {
		if slist[key]< glist[key]{
			B=B+slist[key]
		}else {
			B=B+glist[key]
		}
	}
	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
}
