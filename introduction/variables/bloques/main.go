package main

import (
	"fmt"
)

func main() {

	//BLOQUES > podemos definir variables dentro de un scope, esta usara la memoria solo dentro del scope
	//Todas las variables que metamos dentro de {} viviran en ese scope.
	fmt.Println("----BLOQUES---")
	myScopeVariable1 := 50
	fmt.Println(myScopeVariable1)

	{
		myScopeVariable2 := 40
		fmt.Println(myScopeVariable2)
	}

}
