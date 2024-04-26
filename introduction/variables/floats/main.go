package main

import (
	"fmt"
	"unsafe"
)

func main() {

	/*
		DataType Float, si o si tenemos que definir el espacio que ocupa este valor
		solo tenemos de 32 y 64 bits
	*/
	var myFloat32Var float32 = 3.14
	fmt.Printf("Type: %T, Value: %F, Bytes: %d, Bits: %d,\n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 590.12435
	fmt.Printf("Type: %T, Value: %F, Bytes: %d, Bits: %d,\n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	myOtherFloat := 7654.12 // por defecto tomara 64bits que es lo que tiene mi maquina
	fmt.Printf("Type: %T, Value: %F, Bytes: %d, Bits: %d,\n", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat)*8)
}
