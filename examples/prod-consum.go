// Implement the Producer-Consumer Pattern
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func prod(data chan string, done chan bool) {
	random_string := make([]byte, 10) 

	// Loop para criar uma string aleatória 
	for {
		rand.Seed(time.Now().UnixNano())
		for i:=0; i<10; i++ {
			random_string[i] = byte(65 + rand.Intn(25))
		}
		data <- string(random_string)
		time.Sleep(1*time.Second)
	}
	done <- true
}

func consumer(receive chan string) {
	for {
		fmt.Println(<-receive)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dados := make(chan string)
	done := make(chan bool)

	go prod(dados, done)
	go consumer(dados)

	<-done // similar to join in python
}