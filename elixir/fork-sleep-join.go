//Crie um programa que recebe um número inteiro n como argumento e cria n goroutines.
//Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5
//segundos. A main-goroutine deve esperar todas as goroutines filhas terminarem de
//executar para em seguida escrever na saída padrão o valor de n.

package main

import (
	"fmt"
	"time"
	"math/rand"

)

func rotina(id int, done chan bool) {
	sleep_duration := time.Duration(rand.Intn(5))*time.Second
	fmt.Printf("Eu sou a Go Rotina %d e estou dormindo!\n", id)
	time.Sleep(sleep_duration)
	fmt.Printf("Eu sou a Go Rotina %d e acordei!\n", id)
	done <- true //envia um sinal para a main-goroutine
}

func main() {
	n := 5
	done := make(chan bool)
	
	// cria n goroutines
	for i:=0; i < n; i++ {
		go rotina(i, done)
	}

	// espera todas as goroutines terminarem 
	for i:=0; i < n; i++ {
		<-done
	}

	fmt.Println("Todas as Go Routines terminaram.\nValor de n é:", n)
}
