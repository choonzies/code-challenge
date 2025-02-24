package main

import "fmt"

func sum_to_n_a(n int) int {
	return n * (n + 1) / 2
}

func sum_to_n_b(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sum_to_n_c(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum_to_n_c(n-1)
}

func main() {
	fmt.Println("The first way will be the most optimal way, by using formula n(n+1)/2 derived from sum of AP formula.")
	fmt.Println("Time Complexity: O(1) as this is computed by a simple mathematical formula.")
	fmt.Println("Space Complexity: O(1) as we are not using any extra space.")
	fmt.Println("First way of summing 1 to 10: \n", sum_to_n_a(10))

	fmt.Println("The second way will be the simple solution of a single for loop")
	fmt.Println("Time Complexity: O(n) as we will need to iterate through up till n")
	fmt.Println("Space Complexity: O(1) as we are only using a constant set of additional variables to accumulate the sum.")
	fmt.Println("Second way of summing 1 to 10: \n", sum_to_n_b(10))

	fmt.Println("The third way will be a slightly more complicated recursive solution")
	fmt.Println("Time Complexity: O(n) as we will need to call the function about n times")
	fmt.Println("Space Complexity: O(n) as each recursive call adds a new frame to the function call stack.")
	fmt.Println("Third way of summing 1 to 10: \n", sum_to_n_c(10))
}
