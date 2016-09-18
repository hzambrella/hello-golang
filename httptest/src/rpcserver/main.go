package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

//type Quotient struct{
//	Quo,Rem int
//}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	fmt.Println("rpc ready")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
