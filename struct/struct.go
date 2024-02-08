package main

// Struct se refiere a una coleccion de datos, una variable puede ser almacenada dentro de otra y eso se conoce como un struct.
// Struct puede tener un conjunto de variables, estas estructuras pueden contener metodos tambien, es como una clase en java
//Sigue la misma logica que las funciones, si la comienzan con mayuscula es publico, si la comienzan con minuscula es privado.
/*
	Con la etiqueta de json indicamos como  llamaremos a la variable en el json
*/

import (
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
}
