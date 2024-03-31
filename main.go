package main

import "fmt"

func main() {
	canal := make(chan string)

	//T2
	go func() {
		canal <- "Golang Conference!"
	}()

	//T1
	msg := <-canal
	fmt.Println((msg))
}
