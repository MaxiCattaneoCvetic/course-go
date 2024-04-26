package function

//Dentro del mismo package puedo acceder a todos los recursos

// Ahora decimos que vamos a retornar una funcion que recibe dos float64 y retorna un float64
func FactoryOperation(op Operation) func(float64, float64) float64 {
	switch op {
	case SUM:
		return sum
	case DIV:
		return div
	case MUL:
		return mul
	case SUB:
		return sub
	}
	return nil

}

func sum(a, b float64) float64 {
	return a + b
}

func div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func mul(a, b float64) float64 {
	return a * b
}

func sub(a, b float64) float64 {
	return a - b
}
