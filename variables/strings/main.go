package main

import (
	"fmt"
)

func main() {

	/*
		DataType Float, si o si tenemos que definir el espacio que ocupa este valor
		solo tenemos de 32 y 64 bits
	*/

	var myStringVar3 string = "test1234"
	fmt.Printf("Mi valor %s \n", myStringVar3)

	myStringVar4 := `Mi variable de tipo texto en go 
	con multiples
	lineas
	:)`
	fmt.Printf("%s \n", myStringVar4)
}
