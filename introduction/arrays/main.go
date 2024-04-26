package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var myIntVar int = 30
	fmt.Printf("Type: %T, Value: %d, Bytes: %d, Bits: %d,\n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	//declaramos el array de orden 5
	var myArrayVar1 [5]int
	fmt.Println(myArrayVar1, "- size: ", len(myArrayVar1))

	myArrayVar2 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar2, "- size: ", len(myArrayVar2))

	myArrayVar1[0] = 2
	myArrayVar1[1] = 4
	myArrayVar1[2] = 55
	fmt.Println(myArrayVar1, "- size: ", len(myArrayVar1))

	fmt.Printf("Type: %T, Value: %d, Bytes: %d, Bits: %d,\n", myArrayVar1, myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	// for i := range myArrayVar1 {
	// 	fmt.Println("index: ", i, " value: ", myArrayVar1[i])
	// }

	// el range recibe 2 parametros, el indice y el valor
	for index, value := range myArrayVar1 {
		fmt.Println("index: ", index, " value: ", value)
	}

}
