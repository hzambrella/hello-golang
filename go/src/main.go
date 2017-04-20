package main

import(
	"fmt"
	"bufio"
	"strconv"
	"strings"

	"lib/ipc"
	"ctrl"
	"os"
)

const(
		help=`
		Command:
		i<username><level><exp>         add player
		o<username>				        remove player   
		b<message>                      broadcast      
		l						        list all player
		q                               quit
		h                               help            
	`)

var centerClient *ctrl.CenterClient

func startCenterServer(){

	//centerServer:=ctrl.NewCenterServer()   //wrong
	server:=ipc.NewIpcServer(&ctrl.CenterServer{})
	ipcClient:=ipc.NewIpcClient(server)
	centerClient=&ctrl.CenterClient{ipcClient}
}


func Help(args []string)int{
	fmt.Println(help)
	return 0
}

func Quit(args []string)int{
	fmt.Println("good bye!")
	return 1
}

func Login(args []string)int{
	//in case of panic :index out of range ,should check len(args)
	if len(args)!=4{
		fmt.Println("Usage:i<username><level><exp>")
		return 0
	}

	level,err:=strconv.Atoi(args[2])
	if err!=nil{
		fmt.Println(err.Error())
		panic(err)
		return 0
	}

	exp,err:=strconv.ParseFloat(args[3],64)
	if err!=nil{
		fmt.Println(err.Error())
		panic(err)
		return 0
	}

	p:=&ctrl.Player{Name:args[1],Level:level,Exp:exp}
	if err:=centerClient.AddPlayer(p);err!=nil{
		fmt.Println(err.Error())
		return 0
		//panic(err)
	}
	return 0
}

func Logout(args []string)int{
	//in case of panic :index out of range ,should check len(args)
	if len(args)!=2{
		fmt.Println("Usage:o<username>")
		return 0
	}

	if err:=centerClient.RemovePlayer(args[1]);err!=nil{
		fmt.Println(err.Error())
	}
	return 0
}

func ListPlayer(args []string)int{
	players,err:=centerClient.ListPlayer()
	if err!=nil{
		fmt.Println(err.Error())
		return 0
	}

	for _,v:=range players{
		fmt.Println(fmt.Sprintf("name:%s|lever:%d [exp %.2f]\n",v.Name,v.Level,v.Exp))
	}
	return 0
}

func Broadcast(args []string)int{
	//in case of panic :index out of range ,should check len(args)
	if len(args)!=2{
		fmt.Println("Usage:b<message>")
		return 0
	}
	if err:=centerClient.Broadcast(args[1]);err!=nil{
		fmt.Println(err.Error())
		return 0
	}
	return 0
}

//notice it!
func CommandHandle()map[string]func(args []string)int{
	return map[string]func(args []string)int{
		"i":Login,
		"o":Logout,
		"h":Help,
		"b":Broadcast,
		"l":ListPlayer,
		"q":Quit,
	}
}

func main(){
	startCenterServer()
	handleFunc:=CommandHandle()
	r:=bufio.NewReader(os.Stdin)
	for{
		fmt.Print("Command>>")
		b,_,_:=r.ReadLine()
		token:=strings.Split(string(b)," ")
		if handle,ok:=handleFunc[token[0]];ok{
			ret:=handle(token)
			if ret==1{
				break
			}
		}else{
			fmt.Println(`unknown command`)
			helpFunc:=handleFunc["h"]
			helpFunc(token)
		}
	}
}
