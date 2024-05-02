package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Funcao auxiliar para gerar strings aleatorias
func random_string() string {
	rand.Seed(time.Now().UnixNano()) 
	random_string := make([]byte, 10) 
	
	for i:=0; i<10; i++ {
		random_string[i] = byte(65 + rand.Intn(25))
	}
	word := string(random_string)

	return word
}

// Funçao geradora de strings
func request_stream() chan string {
	ch := make(chan string) 

	go func() {
		for {
			ch <- random_string() 
			time.Sleep(1*time.Second)
		}
	}()
	return ch 
}

// Funcao que ingere strings de um canal 
func ingest(in chan string) {
	for {
		fmt.Println(<- in) // recebe do canal e printa 
	}

}

func main() {
	ch1 := request_stream() // "Me envie strings!!" 
	ch2 := request_stream()

	ch3 := make(chan string) // canal 3 
	go ingest(ch3) // chama uma go routine para ingerir as strings

	go func() {
		for {
			select {
				case s := <- ch1: // envia oq tiver em ch1 para s
					ch3 <- "Canal 1: " + s // envia para ch3 
				case s := <- ch2: 
					ch3 <- "Canal 2: " + s 
			}
		}

	}()
	select {} 

}
