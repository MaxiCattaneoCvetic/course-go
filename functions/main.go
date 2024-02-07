package main

import (
	"GoCourse/functions/function"
	"fmt"
)

func main() {

	// el segundo valor que esta despues de los parametros es el tipo de dato que retornara la function.
	// en el caso de no retornar nada, no podemos nada
	var myIntVar int64
	function.Display(myIntVar)
	function.Suma(2, 3)
	function.RepeatString(5, "Hola")
	function.Calc(function.SUM, 2, 3)
	{
		value, error := function.Calc(function.SUM, 2.12, 34)
		fmt.Println("Valor: ", value, " Error: ", error)
	}
	{
		value, error := function.Calc(function.DIV, 5, 0)
		fmt.Println("Valor: ", value, " Error: ", error)
	}
	{
		value, error := function.Calc(function.TEST, 5, 5)

		if error != nil {
			fmt.Println("Error: ", error)
		} else {
			fmt.Println("Valor: ", value, " Error: ", error)
		}
	}

	xVal, yVal := function.Split(40)
	fmt.Println(xVal, yVal)

	// Podemos tener una funcion que no sepamos que valores recibira, puede recibir cualquier cantidad de parametros
	// para esto usamos los (...)
	val2 := function.MSum(10, 30, 50, 60)
	fmt.Println(val2)

	{
		MoVal, err := function.MOperations(function.DIV)
		fmt.Println(MoVal, err)
	}

	{
		MoVal, err := function.MOperations(function.SUM, 10, 30, 50, 60)
		fmt.Println(MoVal, err)
	}

}
