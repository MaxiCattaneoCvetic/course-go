/*
Con esto podemos pasar disinttos valores a las funciones

*/

package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {

	v1 := []float64{1, 2, 3, 4, 5, 25.25, 20.25, 45}
	v2 := []float32{9, 23, 1, 23, 8, 98}
	//fmt.Println(Sum01(v1))
	//fmt.Println(Sum01(v2))
	fmt.Println(Sum02(v1))
	fmt.Println(Sum02(v2))

	anyType(1, 1)
	anyType("1", "HOLA")
	anyTypeTwo(1, "HOLA")
	comparableType(1, 2)
	orderedValues(3, 4)
	orderedValues2(3, 4)

	// podemos establecer arrays genericos, UN SLICE ES UN ARRAY PERO FLEXIBLE
	csInt := []int{1, 2, 3, 4, 5}
	fmt.Println(csInt)
	csStg := []string{"1", "2", "3", "4", "5"}
	fmt.Println(csStg)

	x, y := Point(5), Point(2)
	fmt.Println(MinNumber(x, y))

	fd := MyGenericStruct[MyFirstData]{StrValue: "test", Data: MyFirstData{}}
	fs := MyGenericStruct[MySecondData]{StrValue: "test", Data: MySecondData{}}

	fd.Data.PrintOne()
	fs.Data.PrintTwo()
}

// generamos una funcion que permita sumar cosas que son distintas, en este caso float 64 + float32
// N es el nombre de generico
func Sum01[N int | int32 | int64 | float32 | float64](v []N) N {

	var total N
	for _, tV := range v {
		total += tV
	}
	return total

}

// puedo generar una interfaz tambien, que este  reprensentada por distintos tipos de datos

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum02[N Number](v []N) N {

	var total N
	for _, tV := range v {
		total += tV
	}
	return total

}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyTypeTwo[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2) // para utilizar esto y ver si son comparables debo poner comparable que es el tipo de generic

}

func orderedValues[N constraints.Ordered](v1, v2 N) {
	fmt.Println(v1, v2, "orders---")

	// Para comparar si es mayo o menor necesitamos un package external
	// go get golang.org/x/exp
	fmt.Println(v1 > v2)
	fmt.Println(v1 < v2)

}

func orderedValues2[N cmp.Ordered](v1, v2 N) { // esta es la funcion nativa de GO sin necesidad de instalar el package
	fmt.Println(v1, v2, "orders---")

	// Para comparar si es mayo o menor necesitamos un package external
	// go get golang.org/x/exp
	fmt.Println(v1 > v2)
	fmt.Println(v1 < v2)

}

type CustomSlice[V int | string] []V

func MinNumber[T N1](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Point int

// en este caso como vamos a mandarle una representacion de un int -> como un POINT le especificamos a la interfaz con ~
type N1 interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type (
	MyGenericStruct[D any] struct {
		StrValue string
		Data     D
	}

	MyFirstData  struct{}
	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print First")
}
func (d MySecondData) PrintTwo() {
	fmt.Println("Print Second")
}
