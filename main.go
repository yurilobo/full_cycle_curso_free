package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", workerID, x)

	}

}

func main() {
	ch := make(chan int) //cria o canal
	go worker(1, ch)     //inicia  o processador
	//go worker(2, ch)        //dois processamentos
	qtdWorkers := 10

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 15; i++ { //procesa o canal
		ch <- i
	}
}

//nesse caso temos workers trabalhando de forma paralelas em cores diferentes da maquina
