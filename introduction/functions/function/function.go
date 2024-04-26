package function

// Si la funcion comienza con letra mayuscula es publica, eso quiere decir que puede ser usada por otros modulos.
// De lo contrario si la funcion comienza con letra miniscula es privada, es decir solo puede ser usada por la misma modulo.
// un modulo es una carpeta que contiene un conjunto de archivos.
import (
	"errors"
	"fmt"
)

type Operation int // tipo de dato para las operaciones

// Definimos constantes con el tipo de dato Operation para las operaciones.
// Generamos un ENUM
const (
	SUM Operation = iota
	SUB
	DIV
	MUL
	TEST
)

func Display(myValue int64) {
	fmt.Printf("Type: %T,value: %d,address: %v \n", myValue, myValue, &myValue)
}

func Suma(a int, b int) int {
	fmt.Println("La suma es: ", a+b)
	return a + b
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

// definimos una funcion que toma como parametro un tipo de dato Operation y dos float64.
// GO nos permite retornar multiples valores, para eso es el segundo parametro de la funcion
func Calc(op Operation, x float64, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("Y must be different than 0")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}

	return 0, errors.New("Invalid operation")
}

// podemos hacer un retorno nombrado, definimos datos para retornar
func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return // en este caso retornaremos ( x e y ) en ese orden
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MSum2(palabra string, numero int, values ...float64) float64 {
	fmt.Println(palabra, numero) // podemos usar la variable palabra y numero y despues recibir un array de float, siempre debe ir al final
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("No values")
	}
	sum := values[0]
	for _, v := range values[1:] {
		switch op {
		case SUM:
			sum += v
		case SUB:
			sum -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("Y must be different than 0")
			} else {
				sum /= v
			}
		case MUL:
			sum *= v
		}
	}
	return sum, nil
}

func MOperationsPRO(op Operation, x, y float64) (float64, error) {
	value, err := Calc(op, x, y)
	fmt.Println("Value: ", value, " Error: ", err)
	return value, err
}
