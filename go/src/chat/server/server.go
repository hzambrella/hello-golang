package main

import (
	"chat/proto"
	"fmt"
	"net"
	"strconv"

	"github.com/hzambrella/gotool/loghz"
)

var logh = loghz.NewLogDebug(true)
var userStore = make([]*proto.User, 0)
var serverName string = "server777"

const (
	setNameMess = "请设置用户name"
)

func handleConnection(conn net.Conn, talkChan map[int]chan string) {
	//fmt.Printf("%p\n", talkChan)  //用以检查是否是传过来的指针

	/*
	   定义当前用户的uid
	*/
	var curUid int

	var err error

	/*
	   定义关闭通道
	*/
	var closed = make(chan bool)

	defer func() {
		logh.Println("defer do : conn closed")
		conn.Close()
		logh.Println("delete userid [%v] from talkChan", curUid)
		delete(talkChan, curUid)
	}()

	/**
	 * 提示用户设置自己的uid， 如果没设置，则不朝下执行
	 */
	for {
		/*
			//提示客户端设置用户id
			_, err = conn.Write([]byte("请设置用户uid"))
			if err != nil {
				logh.Error(err, "设置用户id")
				return
			}
			data := make([]byte, 1024)
			c, err := conn.Read(data)
			if err != nil {
				//closed <- true  //这样会阻塞 | 后面取closed的for循环，没有执行到。
				logh.Error(err, "读取用户id")
				return
			}
			sUid := string(data[0:c])

			//转成int类型
			uid, _ := strconv.Atoi(sUid)
			if uid < 1 {
				continue
			}
		*/
		uid := len(userStore)
		suid := strconv.Itoa(uid)
		curUid = uid
		talkChan[uid] = make(chan string)
		//fmt.Println(conn, "have set uid ", uid, "can talk")
		b, err := userStore[0].MakeMess(proto.SetName, suid, setNameMess)
		if err != nil {
			logh.Error(err, "设置用户name")
			return
		}

		_, err = conn.Write(b)
		if err != nil {
			logh.Error(err, "设置用户name")
			return
		}

		data2 := make([]byte, 1024)
		c, err := conn.Read(data2)
		if err != nil {
			//closed <- true  //这样会阻塞 | 后面取closed的for循环，没有执行到。
			logh.Error(err, "读取用户name")
			return
		}

		nameResp, err := userStore[0].GetMess(data2[0:c])
		if err != nil {
			logh.Error(err, "设置用户name")
			return
		}

		if nameResp.Type != proto.SetName {
			nameResp.Type = proto.SetName
			nameResp.UidTo = proto.ServerUid
		}

		userName := nameResp.Content
		if err != nil {
			logh.Error(err)
			return
		}
		user, err := proto.NewUser(suid, userName)
		userStore = append(userStore, user)

		setNameInfo := "have set uid " + suid + ",name " + userName + " can talk"
		bnameResp, err := userStore[0].MakeMess(proto.Default, suid, setNameInfo)
		if err != nil {
			logh.Error(err)
			return
		}

		_, err = conn.Write(bnameResp)
		if err != nil {
			logh.Error(err, "通知用户可以说话")
			return
		}
		break
	}

	//	logh.Println("err 3")

	//当前所有的连接
	logh.Println(talkChan)

	//读取客户端传过来的数据
	go func() {
		for {
			//不停的读客户端传过来的数据
			fmt.Println("listening ..")
			socket := make([]byte, 1024)
			c, err := conn.Read(socket)
			if err != nil {
				logh.Error(err, "have no client write", "读客户端传过来的数据")
				return
				//closed <- true //这里可以使用 | 因为是用用的go 新开的线程去处理的。 |  即便chan阻塞，后面的也会执行去读 closed 这个chan
			}
			user := userStore[curUid]
			mess, err := user.GetMess(socket[0:c])
			if err != nil {
				logh.Error(err, "读客户端传过来的数据")
				continue
				//closed <- true
			}

			switch mess.Type {
			case proto.Default:
				uto, err := strconv.Atoi(mess.UidTo)
				if err != nil {
					logh.Error(err, "string uid to uid")
					break
				}
				logh.Println("send to:", uto)
				talkChan[uto] <- string(socket[0:c])
			case proto.Public:
				logh.Println("public")
				for k, _ := range talkChan {
					if k != 0 {
						talkChan[k] <- string(socket[0:c])
						logh.Println("send to:", k)
					}
				}
			case proto.SetName:
				logh.Println("invalid type")
				//talkChan[0] <- string(socket[0:c])
			default:
				logh.Println("invalid type")
				//do nothing
			}

			//将客户端过来的数据，写到相应的chan里
			/*
				if curUid == 3 {
					talkChan[4] <- clientString
				} else {
					talkChan[3] <- clientString
				}
			*/
			//talkChan[curUid] <- clientString

		}
	}()

	/*
	   从chan 里读出给这个客户端的数据 然后写到该客户端里
	*/
	go func() {
		for {
			socketStr := <-talkChan[curUid]
			logh.Println(curUid, "is recieve message")
			_, err = conn.Write([]byte(socketStr))
			if err != nil {
				logh.Error(err, "读出给这个客户端的数据 然后写到该客户端里")
				closed <- true
			}
		}
	}()

	/*
	   检查是否已经关闭连接 如果关闭则推出该线程  去执行defer语句
	*/
	for {
		if <-closed {
			logh.Println("end<-close")
			return
		}
	}
}

func main() {

	/**
	  建立监听链接
	*/
	ln, err := net.Listen("tcp", "127.0.0.1:6010")
	if err != nil {
		panic(err)
	}

	//创建一个管道

	//talkChan := map[f]
	talkChan := make(map[int]chan string)

	server, err := proto.NewUser(strconv.Itoa(len(talkChan)), serverName)
	if err != nil {
		logh.Error(err, "服务号建立出错")
		return
	}

	userStore = append(userStore, server)
	talkChan[0] = make(chan string)

	/*
	   监听是否有客户端过来的连接请求
	*/
	for {
		fmt.Println("wait connect...")
		conn, err := ln.Accept()
		if err != nil {
			logh.Error(err, "get client connection error: ")
		}

		go handleConnection(conn, talkChan)
	}
}
