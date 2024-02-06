package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//CONVER TO STRING
	{
		floatVar := 33.11
		fmt.Printf("Type: %T, Value: %f \n", floatVar, floatVar)

		/*
			En resumen, el código toma un número de punto flotante (floatVar), lo
			convierte a una cadena formateada (floatStrVar), y luego imprime
			información sobre esta cadena formateada, mostrando su tipo y valor en la consola
		*/

		floatStrVar := fmt.Sprintf("%.2f", floatVar) // en lugar de imprimirlo por consola, le asignamos ese valor a la variable
		fmt.Printf("Type: %T, Value: %s \n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("Type: %T, Value: %d\n", intVar, intVar)

		fmt.Println(strings.Repeat("-", 20))

		intStrVar := fmt.Sprintf("%d", intVar)
		// convierto a string
		fmt.Printf("Type: %T, Value: %s\n", intStrVar, intStrVar)
		{
			//Otra forma de convertir a string
			fmt.Println("---------convert to string 2 -----------")
			intVar2 := 342
			intStrVar2 := strconv.Itoa(intVar2)
			fmt.Printf("Type: %T, Value: %s\n", intStrVar2, intStrVar2)
		}
	}

	//CONVER TO NUMBER
	{
		fmt.Println("------------Convert string to number-----------")
		strIntVar, error := strconv.Atoi("15") // convierto de string a int

		fmt.Printf("Type: %T, Value: %d Error: %v \n", strIntVar, strIntVar, error)
		// ParseInt recibe 3 parametros, el string, el base y el tamanio de la variable que queremos convertir a int

		/* si es base 10, eso quiere decir que es un numero entero,
		si es base 16, es un numero hexadecimal si es base 8, es un numero octal */

		strIntVar2, _ := strconv.ParseInt("22", 10, 64)              // convierto de string a int
		fmt.Printf("Type: %T, Value: %d \n", strIntVar2, strIntVar2) // -> con el guion bajo, omitimos el parametro, que en este caso es el error
	}

	//CONVER TO FLOAT
	{
		fmt.Println("------------Convert string to float-----------")
		strFloatVar, err := strconv.ParseFloat("3.14", 64) // convierto de string a float, no le paso la base en este caso, sino que le paso el tamaño
		fmt.Printf("Type: %T, Value: %0.3f Error: %v \n", strFloatVar, strFloatVar, err)

	}
}
