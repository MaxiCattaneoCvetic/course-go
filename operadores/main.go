package main

import (
	"fmt"
)

func main() {

	yearsOld := 33

	fmt.Println(!(yearsOld < 40))

	if yearsOld > 18 {
		fmt.Printf("%d es mayor a 18 \n", yearsOld)
	} else {
		fmt.Printf("%d es menor a 18 \n", yearsOld)

	}

	functionSum := func(x, y int) int {
		return x + y
	}

	if value := functionSum(2, 3) >= 5; value {
		fmt.Println("La suma es mayor a 5")
	} else {
		fmt.Println("La suma es menor a 5")
	}

	number := 6

	if number == 1 {
		fmt.Println("El numero es 1")
	} else if number == 2 {
		fmt.Println("El numero es 2")
	} else if number == 3 {
		fmt.Println("El numero es 3")
	} else {
		fmt.Println("El numero no es 1, 2 o 3")
	}

	// Switch
	{
		switch number {
		case 1:
			fmt.Println("El numero es 1")
		case 2:
			fmt.Println("El numero es 2")
		case 3:
			fmt.Println("El numero es 3")
		default:
			fmt.Println("El numero no es 1, 2 o 3")
		}
	}

	// podemos hacer el mismo caso asignando un valor y validando
	switch number := 1; number {
	case 1:
		fmt.Println("El numero es 1")
	case 2:
		fmt.Println("El numero es 2")
	case 3:
		fmt.Println("El numero es 3")
	default:
		fmt.Println("El numero no es 1, 2 o 3")
	}

	// En el caso de que no entre en ningun case, no se ejecuta ya que no tiene el case default
	switch {
	case number > 0:
		fmt.Println("El numero es positivo")
	case number < 9:
		fmt.Println("El numero es negativo")
	case number == 0:
		fmt.Println("El numero es 0")
	}

}
