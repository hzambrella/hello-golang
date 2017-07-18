package main

import (
	"chat/proto"
	"fmt"
	"math/rand"
	"net"

	"github.com/hzambrella/gotool/loghz"
)

var logh = loghz.NewLogDebug(false)
var userClient = &proto.User{Uid: "-1"}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6010")
	if err != nil {
		panic(err)
	}

	logh.Println(&conn)
	fmt.Println("hello server")

	defer conn.Close()
	go writeFromServer(conn)

	for {
		fmt.Println("input:")
		var talkContent string
		fmt.Scanln(&talkContent)
		if len(userClient.Name) == 0 {
			userClient.Name = talkContent
		}
		var uidTo string
		var messType int
		for {
			fmt.Println("input messType 1 :public  0 :default 2 setName: ")
			fmt.Scanln(&messType)
			if messType >= proto.Default || messType <= proto.SetName {
				break
			} else {
				fmt.Println("wrong type ,try again")
			}
		}

		switch messType {
		case proto.SetName:
			uidTo = proto.ServerUid
		case proto.Public:
			uidTo = proto.PublicUid
		default:
			fmt.Println("input uidTo:")
			fmt.Scanln(&uidTo)
		}

		if len(talkContent) > 0 && len(uidTo) > 0 {
			b, err := userClient.MakeMess(messType, uidTo, talkContent)
			//logh.Println(string(b))
			if err != nil {
				fmt.Println("encode error")
				return
			}

			_, err = conn.Write(b)
			if err != nil {
				fmt.Println("write to server error")
				return
			}
		}
	}
}

/*
func connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:6010")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "hello server\n")

	defer conn.Close()
	go writeFromServer(conn)

	for {
		var uid string
		var messType int
		var talkContent string
		fmt.Scanln(&messType)
		fmt.Scanln(&uid)
		fmt.Scanln(&talkContent)

		if len(talkContent) > 0 && len(uid) > 0 {
			proto.MakeMess(messType, uid, talkContent)
			_, err = conn.Write([]byte(talkContent))
			if err != nil {
				logh.Println("write to server error")
				return
			}
		}
	}
}
*/

func writeFromServer(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		c, err := conn.Read(data)
		if err != nil {
			logh.Error(err, "rand", rand.Intn(10), "have no server write")
			return
		}

		mess, err := userClient.GetMess(data[0:c])
		if err != nil {
			logh.Println(err.Error())
			return
		}

		if mess.Type == proto.SetName {
			userClient.Uid = mess.UidTo
		}
		fmt.Println(userClient.Uid, ":recieve:")
		fmt.Println(mess)
		//logh.Println(string(data[0:c]) + "\n ")
	}
}
