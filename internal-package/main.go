package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	p := fmt.Println
	p("------------------PACKAGE TIME------------------")
	now := time.Now()
	p(now)

	then := time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	then = then.Add(1 * time.Hour) // le agregamos 1 hora
	p(then)

	then = then.Add(24 * time.Hour) // le agregamos 1 dia
	p(then)

	p(then.Year())  // obtenemos el dia
	p(then.Month()) // obtenemos el mes
	p(then.Day())   // obtenemos el anio

	p(now.Weekday()) // obtenemos el dia de la semana

	p("Then es antes que Now:", then.Before(now))
	p("Then es despues que now:", then.After(now))
	p("then es igual a now:", then.Equal(now))

	diff := now.Sub(then)
	//dif es la diferencia entre now y then
	p("Diferencia entre now y then: ", diff)

	p("------------------PACKAGE OS------------------")
	// Lo que haces es utilizar las funcionalidades del sistema operativo

	file, err := os.Open("main.go") // declaramos una variable de tipo file y una variable de tipo error
	if err != nil {
		p(err)
		os.Exit(1) // 1 > hay erro 0 > no hay error
	}
	p(file)

	v, _ := file.Stat() // _ es el error y stat() es la variable que obtiene la informacion del archivo
	p(v.Name(), v.Size())

	//data := make([]byte, 200)

	// recibe 2 valores, un numero y un errro
	// c, err := file.Read(data)
	// if err != nil {
	// 	p(err)
	// 	os.Exit(1)
	// }

	//p(string(data[:c]), c)

	p(os.Getenv("JAVA_HOME"))
	p(os.Getenv("USERNAME"))
	p(os.Setenv("asd", "asd"))

}
