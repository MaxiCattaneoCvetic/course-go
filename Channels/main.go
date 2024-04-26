package main

import (
	"fmt"
	"time"
)

func main() {

	// con chan defino que esto sera un canal

	myFirstChannel := make(chan string)
	go myProcessWithChannel("C", myFirstChannel)
	result := <-myFirstChannel // leo el canal y almaceno la variable que envio en el canal
	fmt.Println(result)
	// en la instancia del result quedamos esperando a que el canal tenga un valor, y cuando lo tenga sigo con la ejecución
	close(myFirstChannel) // cierrro el canal cuando termina

	channelA := make(chan string)
	channelB := make(chan string)

	fmt.Println("------------------------------")
	go myProcessWithChannel("A", channelA)
	go myProcessWithChannel("B", channelB)
	resultA := <-channelA
	fmt.Println("D: ", resultA)
	resultB := <-channelB
	fmt.Println("E", resultB)

	close(channelA)
	close(channelB)

}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("procces: %s - num: %d \n", p, i)
	}
	c <- "ok" // asigno un valor al canal

}

func myProcces(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("procces: %s - num: %d \n", p, i)
	}
}
