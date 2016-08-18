package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type Args struct {
	A, B int
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("input :  ./main num.a num.b")
		os.Exit(1)
	}
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	aaa := &Args{a, b}
	var reply int
	err = client.Call("Arith.Multiply", aaa, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
