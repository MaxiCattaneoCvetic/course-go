package main

import (
	"GoCourse/interface/vehicles"
	"fmt"
)

// podemos utilizar polimorfismo con interfaces
func main() {
	Display("Hello")
	Display(123)
	Display(true)

	// definimos un array de strings
	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK", "CAR"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vehicles: %s \n", v) // print del vehicle en el array

		veh, err := vehicles.New(v, 400) // esto devuelve un vehiculo y un error

		if err != nil { // si el error es diferente de nil entonces imprimimos el error
			fmt.Println("Error: ", err)
			fmt.Println()
			continue // continua con el siguiente elemento, omite el resto de la iteracion
		}

		distance := veh.Distance() // distancia es la distancia del vehiculo
		fmt.Printf("Distance:  %2.f", distance)
		fmt.Println()
		d += distance
	}
	fmt.Printf("Total Distance: %2.f \n", d)

}

// cuando generamos la funcion y le decimos que puede recibir una interface, puede recibir cualquier tipo de dato

func Display(value interface{}) {
	fmt.Println(value)
}
