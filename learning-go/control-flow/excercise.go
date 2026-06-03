package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		num = append(num, rand.Intn(100))
	}

	fmt.Println(num, len(num))

	for _, number := range num {
		if number/2 == 0 {
			fmt.Println("Two!")
		} else if number/3 == 0 {
			fmt.Printf("Three!")
		} else if number/2 == 0 && number/3 == 0 {
			fmt.Printf("Six!")
		} else {
			fmt.Println("Never mind")
		}
	}

	var total int

	for i := 0; i < 10; i++ {
		total = total + i // BUG: the := should be an =
		fmt.Println(total)
	}
}
