package main

import (
	"fmt"
)

func main() {
	/*
		A diferencia de los vectores(arrays), que tienen un tamaño estatico, lo que tienene los slices son dinamicos, es decir, puede cambiar de tamaño.
		Gestionan mejor la momoria que los vectores.
	*/

	myArrayVar := [5]int{3, 6, 9, 10, 16}
	fmt.Println("array: ", myArrayVar, " size: ", len(myArrayVar))

	//Si queremos modificar este array en tiempo de ejecucion no podemos hacerlo, para eso usamos los slices.

	mySliceVar := []int{}
	mySliceVar = append(mySliceVar, 12, 34, 54)
	fmt.Println("array: ", mySliceVar, " size: ", len(mySliceVar))

	{
		mySliceVar2 := myArrayVar[2:4] // myArrayVar[:4] podriamos utilizar para tomar del 0 al 4 o sino myArrayVar[2:] --> tomariamos del 2 al final
		fmt.Println("array: ", mySliceVar2, " size: ", len(mySliceVar2))
	}

	{
		mysliceVar4 := make([]int, 5)
		fmt.Println("array: ", mysliceVar4, " size: ", len(mysliceVar4))
	}

}
