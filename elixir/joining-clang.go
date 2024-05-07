// Considere a API abaixo. A função gateway deve criar e iniciar nthreads (go routines)
// pthreads diferentes. O código executado por cada pthread deve ser o da função
// request. A função request deve sortear um número aleatório n e dormir n segundos.
// Após criar as pthreads, a função gateway deve esperar que até wait_nthreads
// terminem. Após a espera, a função gateway deve retornar a soma dos valores n
// sorteados nas funções request.
// int gateway(int nthreads, int wait_nthreads)
// void* request(void*)

package main

import (
	"fmt"
	"time"
	"math/rand"
)

// A função gateway deve criar e iniciar nthreads (go routines) diferentes.
// O código executado por cada pthread deve ser o da função request.
// Após criar as pthreads, a função gateway deve esperar que até wait_nthreads
// terminem. Após a espera, a função gateway deve retornar a soma dos valores n
func gateway(nGoroutines int) int {
	done := make(chan bool)
	var sumRequests int

	for i:=0; i<nGoroutines; i++ {
		go func(id int) {
			num_req := request()
			sumRequests += num_req 
			fmt.Printf("GoRoutine ID: %d \tNumero sorteado da funcao request(): %d \tSoma atual: %d \n", id+1, num_req, sumRequests) 
			done <- true
		}(i)
	}
	for i:=0; i<nGoroutines; i++ {
		<- done
	}
	return sumRequests
} 

//A função request deve sortear um número aleatório n e dormir n segundos.
func request() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(99)
	sleepDuration := time.Duration(n) * time.Second
	time.Sleep(sleepDuration)
	return n 
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nGoroutines := rand.Intn(99)
	
	fmt.Printf("***--- NÚMERO DE GOROUTINES SORTEADO: %d ---* \n\n", nGoroutines)

	fmt.Println("\nTodas as Go Routines terminaram. Soma dos valores sorteados: ", gateway(nGoroutines))
	
}	