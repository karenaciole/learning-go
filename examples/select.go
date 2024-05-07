package main 

import (
	"fmt"
	"time"
	"math/rand"
)


func sendNumber(ch chan int, n int) {
	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(100)
	// adicionar um time.sleep pode mudar a ordem do select
	ch <- n
}

func sendString(ch chan string, s string) {
	rand.Seed(time.Now().UnixNano())
	chanCharset := make([]byte, 10)
	for i := range chanCharset {
		chanCharset[i] = byte(65 + rand.Intn(25))
	}
	s = string(chanCharset)
	// adicionar um time.sleep pode mudar a ordem do select
	ch <- s
}

func main() {
	chNumber := make(chan int)
	chString := make(chan string)

	go sendNumber(chNumber, 10)
	go sendString(chString, "Golang supremacy")
	
	select {
		// the select is a non deterministic statement
	case n := <-chNumber:
		fmt.Println("Received number:", n)
	
	case s := <-chString:
		fmt.Println("Received string:", s)
	}
	
}