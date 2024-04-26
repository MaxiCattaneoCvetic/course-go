// hacemos referencia a un valor en memoria especifico -> una direccion en memoria

package main

import "fmt"

func main() {
	//declaracion de un puntero

	var i int
	var iP *int
	// definimos una variable i y definimos otra variable IP que almacena la direccion de memoria de i -> puntero
	i = 34
	iP = &i

	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	// modificamos el puntero
	*iP = 1

	fmt.Printf("value i: %d, value pointer i: %d\n", i, *iP)

	myVar := 30
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	increment(myVar)
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	fmt.Println("ejecutando incrementP")
	incrementP(&myVar)
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	fmt.Println("-----------------------------------------")
	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("addres: %p, value: %v \n", mySlice, mySlice)
	fmt.Printf("addr1: %p, addr2: %p, addr2: %p \n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println("mySlice: ", mySlice)

	fmt.Println("---------------Nueva estructura----------------")

	// generamos un puntero de tipo MyStruct
	MyStruct := &MyStruct{
		ID:   1212,
		Name: "test",
	}
	fmt.Printf("address: %p \n", MyStruct)
	fmt.Printf("ID: %d, name: %s \n", MyStruct.ID, MyStruct.Name)
	updateStruct(MyStruct)
	fmt.Printf("id: %d, name: %s \n", MyStruct.ID, MyStruct.Name)

	fmt.Println("---------------test----------------")
	MyStruct.Set("test method")
	fmt.Printf("id: %d, name: %s \n", MyStruct.ID, MyStruct.Name)
	MyStruct.SetP("test method 2")
	fmt.Printf("id: %d, name: %s \n", MyStruct.ID, MyStruct.Name)

}

func increment(val int) {
	val++
	fmt.Println(&val)
	fmt.Println(val)
}

func incrementP(val *int) {
	*val++
	fmt.Println(&val) // direccion de memoria del puntero
	fmt.Println(val)  // valor de la variable que tiene el puntero
	fmt.Println(*val) // donde apunta el puntero

}

func updateSlice(v []int) {
	fmt.Printf("address: %p \n", v)
	fmt.Printf("addr1: %p, addr2: %p, addr2: %p \n", &v[0], &v[1], &v[2])
	v[0] = 12

}

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p \n", &m)
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("address: %p \n", m)
	m.Name = name
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address in function: %p \n", v)
	v.ID = 999
	v.Name = "Update"

}
