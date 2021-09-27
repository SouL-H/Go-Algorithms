package main

import "fmt"

//Dinamik Programlama
func main() {
	n := 15
	memo := make([]int, n)
	fmt.Println(Fib(n))
	fmt.Println(FibOptimzie(n, memo))
	fmt.Println(FibBottomUp(n,memo))
}

func Fib(n int) int {

	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

//Memory solution O(n) karmaşıklığa indirdik.
func FibOptimzie(n int, memo []int) int {
	if n <= 1 {
		return 1
	}
	if memo[n-1] != 0 {
		return memo[n-1]

	}
	res := FibOptimzie(n-1, memo) + FibOptimzie((n-2), memo)
	memo[n-1] = res
	return res
}

func FibBottomUp(n int, k []int) int {
	if n <= 1 {
		return 1
	}
	k[0] = 1
	k[1] = 1

	//i=2 1 1 2
	//i=3 1 1 2 3
	for i := 2; i <= n; i++ {
		k[i] = k[i-1] + k[i-2]
	}
	return k[n]

}
