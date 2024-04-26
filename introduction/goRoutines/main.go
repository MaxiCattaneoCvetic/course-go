/*
Nos permiten ejecutar proceso en distintos hilos de forma simultanea

Ejemplos
proceso 1 > proceso 3
proceso 2 > proceso 4


El proceso 3 depende del proceso 1, cuando este termiune comenzara el proceso 3. Los dos se encuentran bajo el mismo hilo de ejecucion.
Pero al iniciar el proceso general el proceso 1 y 2 se ejecutaran en simultaneo, ya que estan en el mismo tiempo pero en distintos hilos.

LA CANTIDAD DE HILOS DE EJECUCION SON LA CANTIDAD DE NUCLEOS QUE TIENE MI COMPUTADORA

*/

package main

import "fmt"

func main() {

	// en este caso los proceso se ejecutan en mismos hilos

	/*
		myProcces("A")
		myProcces("B")
	*/

	// en este caso los proceso se ejecutan en hilos disnitnos
	go myProcces("A")
	go myProcces("B")
}

func myProcces(p string) {
	i := 0
	for i < 15 {
		i++
		fmt.Printf("procces: %s - num: %d \n", p, i)
	}
}
