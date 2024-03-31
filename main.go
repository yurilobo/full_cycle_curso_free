package main

import "fmt"

func main() {
	ch := make(chan chan int)
	go publish(ch)
	consume(ch)
}
func publish(canal chan int) {
	for i := 0; i < 10; i++ {
		canal <- i
	}
	close(canal)
}

func consume(canal chan int) {
	for x := range canal { //esvazia o canal
		fmt.Println("Mensagem Processada:", x)

	}
	fmt.Println("Todos os itens foram processados")
}
