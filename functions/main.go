package main

import "fmt"

func main() {

	// el segundo valor que esta despues de los parametros es el tipo de dato que retornara la function.
	// en el caso de no retornar nada, no podemos nada
	
	func myFunction(myvalue int64) int64 {
		return myvalue
	}

	fmt.Println(myFunction(12))

}
