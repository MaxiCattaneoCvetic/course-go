package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum++
	}

	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	sum = 0

	for { // El for sin ninguna condicion se ejecuta infinitamente
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)
}
