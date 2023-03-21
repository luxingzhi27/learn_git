package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(sum(3536738, 77, 38))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(3, 3, 322, 22, 444, 99))
}

func sum(args ...int) int {
	ans := 0
	for _, val := range args {
		ans += val
	}
	return ans
}
