package main 

import (
	"fmt"
	"time"
	"math/rand"
)

const maxCapacity = 10


type Request struct {
	id int 
}

var req_ch chan Request =  make(chan Request, maxCapacity)

func create_req() Request {
	random := rand.Intn(100)
	return Request{id : random}
}

func exec_req(req Request) {
	time.Sleep(1*time.Second)
	fmt.Printf("req %d\n", req.id)
}

func worker() {
	for {
		req := <- req_ch
		exec_req(req)	
		time.Sleep(2*time.Second)
		fmt.Println("Recebendo novas requisições...")
	}
}

func main() {
	for i:=0; i<maxCapacity; i++ {	
		go func() {
			for {
				req := create_req() 
				select {
				case req_ch <- req:
				default:
				}
			}
		}()
		go worker()
	}
	select{} 

}
