package main

import (
	"fmt"
	"math/rand"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6010")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "hello server\n")

	defer conn.Close()
	go writeFromServer(conn)

	for {
		var talkContent string
		fmt.Scanln(&talkContent)

		if len(talkContent) > 0 {
			_, err = conn.Write([]byte(talkContent))
			if err != nil {
				fmt.Println("write to server error")
				return
			}
		}
	}
}

func connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:6010")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "hello server\n")

	defer conn.Close()
	go writeFromServer(conn)

	for {
		var talkContent string
		fmt.Scanln(&talkContent)

		if len(talkContent) > 0 {
			_, err = conn.Write([]byte(talkContent))
			if err != nil {
				fmt.Println("write to server error")
				return
			}
		}
	}
}

func writeFromServer(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		c, err := conn.Read(data)
		if err != nil {
			fmt.Println("rand", rand.Intn(10), "have no server write", err)
			return
		}
		fmt.Println(string(data[0:c]) + "\n ")
	}
}
