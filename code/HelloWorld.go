package main

import "fmt"

func main() {

	if _, sub, sum := operations(10, 5); sub > 100 && sum > 20 {
		fmt.Println("The answer is big ", sub)
	} else {
		fmt.Println("The answer is smaller than hundred: ", sum)
	}
}

func operations(x int, y int) (int, int, int) {
	return x * y, x - y, x + y
}
