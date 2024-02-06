package main

import (
	"fmt"
)

func main() {
	/*
		Maps --> Son estructuras representadas como pares clave valor.
	*/

	//Entre [] definimos la Key que va a tener nuestro map -> en este caso sera un int
	map1 := make(map[int]string)

	//Agregamos la key 1
	map1[1] = "A"
	//Agregamos la key 2
	map1[2] = "B"
	map1[333] = "Z"
	fmt.Println(map1)

	map2 := make(map[int][]string)
	map2[1] = []string{"A", "B"}
	map2[3] = []string{"Z", "H", "B"}
	fmt.Println(map2)
	fmt.Println(map2[3])

	// Para modificar un valor de un map
	map1[1] = "New Value"
	fmt.Println(map1[1])

	// Para eliminar un valor de un map -> Le pasamos el map y la key
	delete(map1, 1)

	// ver si existe un valor en map
	valor, existe := map1[1]
	fmt.Println("El valor es: ", valor, " y existe: ", existe)

	map3 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(map3)

	//recorriendo un map
	for key, value := range map3 {
		fmt.Println("key:", key, "- Value:", value)
	}

	//recorriendo map con multiples valores
	// el segundo [] -> indica la cantidad de valores que puede contener cada key
	map4 := map[int][]string{
		1: {"A", "B", "C"},
		2: {"D", "E", "F"},
		3: {"G", "H", "I"},
	}
	fmt.Println(map4)
	fmt.Println("----------------------")
	for key, value := range map4 {
		fmt.Println("key:", key)
		for _, value2 := range value {
			fmt.Println("Value:", value2)
		}

		fmt.Println()
	}

}
