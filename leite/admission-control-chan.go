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
	fmt.Printf("req %d\n", req.id)
	time.Sleep(time.Millisecond*500)
}



func main() {
		go func() {
			for req := range req_ch {
				exec_req(req)
			} 
						
	}()

	for {
		if len(req_ch) < maxCapacity {
			req := create_req()
			req_ch <- req
		} else {
			time.Sleep(1*time.Second)
			fmt.Println("O buffer alcançou a capacidade maxima...")
		}

	}
	

}
