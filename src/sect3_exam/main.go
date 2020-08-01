package main

import "fmt"

func main() {
	test_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range test_slice {
		if r := v % 2; r == 1 {
			fmt.Printf("%v is odd\n", v)
		} else {
			fmt.Printf("%v is even\n", v)
		}
	}
}
