package main

import "fmt"

func main() {

	fmt.Print(numTeams([]int{1, 2, 3, 4}))

}

func numTeams(rating []int) int {
	res := 0
	for i := 0; i < len(rating)-2; i++ {
		for j := i + 1; j < len(rating)-1; j++ {
			if rating[i] > rating[j] {
				for k := j + 1; k < len(rating); k++ {
					if rating[j] > rating[k] {
						res++
					}
				}
			} else {
				for k := j + 1; k < len(rating); k++ {
					if rating[j] < rating[k] {
						res++
					}
				}
			}

		}
	}
	return res

}

// func Section(arr []int) int {
// 	res := 0
// 	for i := 0; i < len(arr)-2; i++ {
// 		for j := i + 1; j < len(arr)-1; j++ {
// 			if arr[i] > arr[j] {
// 				for k := j + 1; k < len(arr); k++ {
// 					if arr[j] > arr[k] {
// 						res++
// 					}
// 				}
// 			} else {
// 				for k := j + 1; k < len(arr); k++ {
// 					if arr[j] < arr[k] {
// 						res++
// 					}
// 				}
// 			}

// 		}

// 	}
// 	fmt.Println()
// 	fmt.Print(res)
// 	return res
// }
