package main

import "fmt"

func main() {

	fmt.Println(countPrimes(499979))
}
func countPrimes(n int) (cnt int) {
	isNotPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			cnt++
			for j := i * i; j < n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return
}