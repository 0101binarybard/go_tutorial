package main

import "fmt"

func main() {

	ar := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}

	for _, v := range ar {
		if v%2 == 0 {
			fmt.Println(v, " is even")
		} else {
			fmt.Println(v, " is odd")
		}
	}
}
