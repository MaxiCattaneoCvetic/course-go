package main

// Struct se refiere a una coleccion de datos, una variable puede ser almacenada dentro de otra y eso se conoce como un struct.
// Struct puede tener un conjunto de variables, estas estructuras pueden contener metodos tambien, es como una clase en java
//Sigue la misma logica que las funciones, si la comienzan con mayuscula es publico, si la comienzan con minuscula es privado.
/*
	Con la etiqueta de json indicamos como  llamaremos a la variable en el json
*/

import (
	"GoCourse/struct/commerce"
	"encoding/json" // Esto nos permite transformar una estructura a json<
	"fmt"
)

// Usamos omitempty para que no aparezca el campo vacio
type User struct {
	ID       int    `json:"id"` //-------> Agregamos json: para especificar que nombre tendra en json
	Name     string `json:"name"`
	Address  string `json:"address,omitempty"`
	Age      int    `json:"age,omitempty"`
	DNI      string `json:"dni,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Gender   string `json:"gender,omitempty"`
}

func main() {
	user := User{
		ID:   1,
		Name: "Cristian",
	}
	fmt.Println(user)

	// Transformamos a json
	v, err := json.Marshal(user)
	fmt.Println(err)
	fmt.Println(v)
	// para hacer la conversion de bites a string  vamos a "castear"
	fmt.Println(string(v))

	/// COMMERCE
	fmt.Println("------------------------------COMMERCE---------------------")
	p1 := commerce.Product{
		Name:  "Heladera marca samsung",
		Price: 2000,
		Type:  commerce.Type{ID: 1, Code: "A", Description: "Electronica"},
		Tags:  []string{"electronic", "freezer", "samsung", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Naranja",
		Price: 50,
		Type:  commerce.Type{ID: 1, Code: "A", Description: "Frutas"},
		Tags:  []string{"alimento", "verdura"},
		Count: 20,
	}

	p3 := commerce.Product{
		Name:  "Cortina",
		Price: 6000,
		Type:  commerce.Type{ID: 1, Code: "C", Description: "Hogar"},
		Tags:  []string{"hogar,cortinas"},
		Count: 3,
	}

	// creamos un carrito
	car := commerce.NewCar(11312)
	car.AddProduct(p1, p2, p3)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total productos: ", len(car.Products))
	fmt.Printf("Total Car %.2f\n", car.Total())
	fmt.Println()

}
