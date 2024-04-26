package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//VARIABLES MODULO 1
	{
		// declaramos la variable
		var myIntVar int
		myIntVar = -12
		fmt.Println("My variable is: ", myIntVar)

		// son igual que los integer, pero positivos ( no acepta valores negativos)
		var myUintVar uint
		myUintVar = 12

		fmt.Println("My variable is: ", myUintVar)

		//String
		var myString string
		myString = "My string variable"
		fmt.Println("My variable is: ", myString)

		//String
		var myBoolean bool
		fmt.Println("My variable is: ", myBoolean)

		// Para ver la ubicacion en memoria de la variable &
		fmt.Println("My variable address is: ", &myBoolean)

		myIntVar2 := 12
		fmt.Println("My variable intValue2 is: ", myIntVar2)

	}

	//CONSTANTES
	{
		fmt.Println("----------------------------------")
		const myIntConst int = 15
		fmt.Println("This is my constant variable: ", myIntConst)
		// Si no le pasamos el tipo de dato, lo guarda como el input que le demos, en este caso String
		const stringConst = "a12"
		fmt.Println("This is my constant variable2: ", stringConst)
	}
	//VARIABLE SIZE -> Podemos declarar la cantidad de bits para cada variable
	fmt.Println("------------------SIZE VARIABLE -------------------")
	{
		// Sizeof() me devuelve el formato en bytes ---> para transformarlo en bits lo multiplico por 8
		var my8BitsUnitVar uint8 = 20
		fmt.Printf("Type: %T, Value: %d, Bytes: %d, Bits: %d\n", my8BitsUnitVar, my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

		var my16BitsUnit uint16 = 30
		fmt.Printf("Type: %T, Value: %d, Bytes: %d, Bits: %d\n", my16BitsUnit, my16BitsUnit, unsafe.Sizeof(my16BitsUnit), unsafe.Sizeof(my16BitsUnit)*8)

		var my32BitsUnit uint32 = 10
		fmt.Printf("Type: %T, Value: %d, Bytes: %d, Bits: %d\n", my32BitsUnit, my32BitsUnit, unsafe.Sizeof(my32BitsUnit), unsafe.Sizeof(my32BitsUnit)*8)

		var my64BitsUnit uint64 = 90
		fmt.Printf("Type: %T, Value: %d, Bytes: %d, Bits: %d\n", my64BitsUnit, my64BitsUnit, unsafe.Sizeof(my64BitsUnit), unsafe.Sizeof(my64BitsUnit)*8)

	}

}
