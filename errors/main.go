// en go podemos manejar los errores.

package main

import (
	"errors"
	"fmt"

	"GoCourse/errors/mypackage"
)

// define una estructura llamada MyCustomError que tiene un campo ID de tipo string
type MyCustomError struct {
	ID string
}

// método Error() para la estructura MyCustomError. Este método retorna un string que describe el error.
// Utiliza fmt.Sprintf para formatear el string con el ID del error.
func (m MyCustomError) Error() string {
	return fmt.Sprintf("Error with id: %s", m.ID) // con SPRINTF retorna un valor string
}

// esta funcion retorna un error con el ID 4
func TestError(v int) error {
	if v == 1 {
		return MyCustomError{"4"}
	} else {
		return errors.New("My other error")
	}
}

func main() {
	// generamos una variable de tipo error
	var err error
	err = errors.New("This is an error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("My format error, string %s integer %d", "Hello", 12)
	fmt.Println(err2)

	err3 := TestError(1)
	fmt.Println(errors.As(err3, &MyCustomError{})) // con el AS chequeamos si el error es de tipo MyCustomError

	err4 := TestError(4)
	fmt.Println(err4)

	err5 := errors.Join(err3, err4)
	fmt.Println("soy el err5:", err5)

	// podemos consultar si un error se encuentra dentro de otro error
	// consultamos si err5 contiene err3 por ejemplo
	fmt.Println(errors.Is(err5, err3))

	// podemos anidar errores
	err6 := fmt.Errorf("error3: [%w] ", err3)
	err7 := fmt.Errorf("errorX: [%w] ", err6)
	fmt.Println("soy el err7:", err7)

	// el Unwrap() permite obtener el error que se encuentra dentro de otro error
	fmt.Println(errors.Unwrap(err7)) // me devuelve el error que fue anidado, osea el err3
	// cuando desenglosamos todos los errores se nos devolvera un nil

	fmt.Println("----------Panic, Defer y Recover-------------")
	/*
		La función defer pospone la ejecución de la función anónima hasta que la función actual termine.
		Esto es útil para tareas como limpieza o manejo de errores al final de una función.
		Cuando hay un panic el control se transfiere automaticamente a la función defer.

		EN GENERAL EL DEFER SE PONE EN EL PRINCIPIO
	*/

	defer func() {
		fmt.Println("end")
		r := recover() // recuperamos el error
		if r != nil {
			fmt.Println("Recovered: ", r)
		}
		// en este caso lanza un error por divivir por 0, entonces agarramos el error con el recover y lo imprimimos
	}()
	// panic("My panic")
	mypackage.Run()
	v := 0
	_ = 5 / v
	fmt.Println("end main")

}
