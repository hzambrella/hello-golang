package httpRpcClient

import(
	"net/rpc"
	"fmt"
)

type Args struct{
	A,B int
}

type Reply struct{
	Result int
}

func ClientRpc(){
	args:=&Args{1,2}
	reply:=new(Reply)
	serverAddress:="localhost"
	client,err:=rpc.DialHTTP("tcp",serverAddress+":1234")
	if err!=nil{
		fmt.Println(err.Error())
	}

	if err:=client.Call("MathRpc.Mutiply",args,&reply);err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(reply.Result)
}
