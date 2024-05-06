//Crie um programa que recebe um número inteiro n como argumento e cria n goroutines.
//Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5
//segundos. A main-goroutine deve esperar todas as goroutines filhas terminarem de
//executar para em seguida escrever na saída padrão o valor de n.

package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"

)

func rotina() {
	fmt.Printf("Eu sou uma Go Rotina e estou dormindo!\n")
}

func main() {
	var wg sync.WaitGroup
	n := 5
	
	for i:=0; i < n; i++ {
		wg.Add(1) // adiciona uma goroutine ao grupo
		go func(id int) {
			defer wg.Done() // remove uma goroutine do grupo
			// cria goroutine filha
			fmt.Printf("ID: %d - ", id)
			go rotina()
			time.Sleep(time.Duration(rand.Intn(5))*time.Second)
			fmt.Printf("ID: %d - Acordei!\n", id)
		}(i)
	}

	wg.Wait() // espera todas as goroutines terminarem
	fmt.Println("Todas as Go Routines terminaram.\nValor de n é:", n)
}
