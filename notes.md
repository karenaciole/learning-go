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
