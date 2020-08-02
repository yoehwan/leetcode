package main

import "fmt"

func main() {
	fmt.Print(buddyStrings("acccccb", "bccccca"))
}
func buddyStrings(A string, B string) bool {
	tmp := make([]int, 26)
	indexs := []int{}
	if len(A) != len(B) { return false }
	if A == B {
		for _, v := range A {
			tmp[v-'a']++
			if tmp[v-'a'] == 2 {
				return true
			}
		}
	}
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			indexs = append(indexs, i)
		}
	}
	if len(indexs) == 2 {
		return A[indexs[0]] == B[indexs[1]] && A[indexs[1]] == B[indexs[0]]
	}
	return  false
}