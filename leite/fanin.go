// Considere a API abaixo como uma função que retorna 
//um canal no qual um número indeterminado de strings serão enviadas.

// func request_stream() chan string

// considere que em um programa, dois canais destes estão sendo manipulados
// da seguinte forma:

// func main() {
// ch1 := request_stream()
// ch2 := request_stream()
// }

// considere que você deve incluir na sua função main uma chamada para a função:

// func ingest(in chan string)

// que deve ser chamada como uma nova goroutine, ou seja:

// go ingest

// agora, o canal recebido pela função ingest deve conter itens disponibilizados 
//pelos canais ch1 e ch2 na medida em que estiverem disponíveis. 
//Ou seja, tão logo itens de ch1 e ch2 estejam disponíveis, estes podem ser enviados para o canal a ser passado.

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
	select {} // o select {} é uma forma de manter a go routine rodando para sempre 

}
