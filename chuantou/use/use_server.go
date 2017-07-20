package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

var localPort string = "53002"
var remotePort string = "60012"

type client struct {
	conn     net.Conn
	er       chan bool
	heart    chan bool
	disheart bool
	writ     chan bool
	recv     chan []byte
	send     chan []byte
}

func (self *client) read() {
	for {
		self.conn.SetReadDeadline(time.Now().Add(time.Second * 40))
		var recv []byte = make([]byte, 10240)
		n, err := self.conn.Read(recv)

		if err != nil {
			self.heart <- true
			self.er <- true
			self.writ <- true
		}
		if recv[0] == 'h' && recv[1] == 'h' {
			self.conn.Write([]byte("hh"))
			continue
		}
		self.recv <- recv[:n]

	}
}

func (self client) write() {

	for {
		var send []byte = make([]byte, 10240)
		select {
		case send = <-self.send:
			self.conn.Write(send)
		case <-self.writ:
			break

		}
	}

}

type user struct {
	conn net.Conn
	er   chan bool
	writ chan bool
	recv chan []byte
	send chan []byte
}

func (self user) read() {
	self.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 800))
	for {
		var recv []byte = make([]byte, 10240)
		n, err := self.conn.Read(recv)
		self.conn.SetReadDeadline(time.Time{})
		if err != nil {

			self.er <- true
			self.writ <- true
			break
		}
		self.recv <- recv[:n]
	}
}

func (self user) write() {

	for {
		var send []byte = make([]byte, 10240)
		select {
		case send = <-self.send:
			self.conn.Write(send)
		case <-self.writ:
			break

		}
	}

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	local, _ := strconv.Atoi(localPort)
	remote, _ := strconv.Atoi(remotePort)
	if !(local >= 0 && local < 65536) {
		fmt.Println("端口设置错误")
		os.Exit(1)
	}
	if !(remote >= 0 && remote < 65536) {
		fmt.Println("端口设置错误")
		os.Exit(1)
	}

	c, err := net.Listen("tcp", ":"+remotePort)
	log(err)
	u, err := net.Listen("tcp", ":"+localPort)
	log(err)
TOP:
	Uconn := make(chan net.Conn)
	go goaccept(u, Uconn)
	fmt.Println("准备好连接了")
	clientconnn := accept(c)
	fmt.Println("client已连接", clientconnn.LocalAddr().String())
	recv := make(chan []byte)
	send := make(chan []byte)
	heart := make(chan bool, 1)
	er := make(chan bool, 1)
	writ := make(chan bool)
	client := &client{clientconnn, er, heart, false, writ, recv, send}
	go client.read()
	go client.write()

	for {
		select {
		case <-client.heart:
			goto TOP
		case userconnn := <-Uconn:
			fmt.Println("user:", userconnn)
			client.disheart = true
			recv = make(chan []byte)
			send = make(chan []byte)
			er = make(chan bool, 1)
			writ = make(chan bool)
			user := &user{userconnn, er, writ, recv, send}
			go user.read()
			go user.write()
			go handle(client, user)
			goto TOP
		}

	}

}

func accept(con net.Listener) net.Conn {
	CorU, err := con.Accept()
	logExit(err)
	return CorU
}

func goaccept(con net.Listener, Uconn chan net.Conn) {
	CorU, err := con.Accept()
	logExit(err)
	Uconn <- CorU
}

func log(err error) {
	if err != nil {
		fmt.Printf("出现错误： %v\n", err)
	}
}

func logExit(err error) {
	if err != nil {
		runtime.Goexit()
	}
}

func logClose(err error, conn net.Conn) {
	if err != nil {
		runtime.Goexit()
	}
}

func handle(client *client, user *user) {
	for {
		var clientrecv = make([]byte, 10240)
		var userrecv = make([]byte, 10240)
		select {

		case clientrecv = <-client.recv:
			user.send <- clientrecv
		case userrecv = <-user.recv:
			client.send <- userrecv
		case <-user.er:
			user.conn.Close()
			runtime.Goexit()
		case <-client.er:
			user.conn.Close()
			client.conn.Close()
			runtime.Goexit()
		}
	}
}
