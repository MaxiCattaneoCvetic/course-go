package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background() //Nos devuelve un contexto
	/*
		podemos mandarle diferentes valores, por ejemplo a un servicio, controller etc
	*/

	// recibe el contexto, una key y un valor
	ctx = context.WithValue(ctx, "key", "my value")
	ctx = context.WithValue(ctx, "key-int", 5)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // cuando termina el main que se ejecute el defer

	// Esperamos hasta que el contexto nos devuelva la info que se exeda  el tiempo

	// almacenamos el contexto en una variable
	go myProcess(ctx2)

	select {
	case <-ctx2.Done():
		fmt.Println("We've timed out")
		fmt.Println(ctx.Err())
	}

}

func viewContext(ctx context.Context) {
	// creamos una funcion que recibe un contexto, y nos devuelve el valor de la key
	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx.Value("key-int"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Si el context  termina, imprimimos por pantalla
			fmt.Println("ohh, time out")
			return
		default:
			fmt.Println("still working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
