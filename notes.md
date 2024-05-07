# Anotações

Sumário
1. [For Loop](#for-loop)
2. [Channel](#Channel)

## For Loop
_Em Go, só existe For como estrutura de loop._ 

### Range

```
package main
import "fmt"

func main() {
	banda := []string{"Kurt", "Dave", "Krist"} // criando um array com dados
	
	for _, integrante := range banda { // use "_" caso vc nao precise do index
		fmt.Println("Integrante: ", integrante)
	}
}
``` 
Saída: 
```
Integrante: Kurt
Integrante: Dave
Integrante: Krist
```

### For infinito 
```
package main

import (
        "fmt"
        "time"
)

func main() {
        i:=0
        for {
                fmt.Println(i)
                time.Sleep(1*time.Second) // tava muito rapido
                i+=1
        }
}
``` 
Saída: 
```
0
1
2
3
4
```

### For com time.Tick
_Você pode usar quando precisa fazer alguma operação dentro do loop a cada X tempo_ 

## Channel
- _Um canal nada mais é que um meio de comunicação entre goroutines. Podemos dizer também que ele atua como um buffer._ 

- _Um canal pode receber e enviar dados. Ele pode ser bidirecional (recebe e envia dados) ou pode ser unidirecional._

- _É possível declarar um canal de tamanho fixo, ou seja, ele só aceitará um número limitado de dados._
```
package main

func main() {
        ch := make(chan int, 1) // criando um canal de tamanho 1
        ch <- 42 // enviando um dado para o canal
        ch <- 27 
        fmt.Println(<-ch) // recebendo um dado do canal

}
```
Saída: 
```	
fatal error: all goroutines are asleep - deadlock!
```
Isso acontece porque o canal só aceita um dado, e quando tentamos enviar o segundo, ele trava.

