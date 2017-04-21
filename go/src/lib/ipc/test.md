/*
package ipc
import(
	"testing"
	"fmt"
)

type EchoServer struct{
}

func (server *EchoServer)Handle(method,params string)*Response{
	return &Response{"200",method+" "+params}

}

func (server *EchoServer)Name()string{
	return "EchoServer"
}

func TestIpc(t *testing.T){
	server:=NewIpcServer(&EchoServer{})
	client1:=NewIpcClient(server)
	client2:=NewIpcClient(server)
	resp1,err:=client1.Call("Caonima","From Client1")
	if err!=nil{
		t.Fatal(err)
	}

	resp2,err:=client1.Call("Caonima","From Client2")
	if err!=nil{
		t.Fatal(err)
	}

	if resp1.Body!="Caonima From Client1"||resp2.Body!="Caonima From Client2"||resp1.Code!="200"||resp2.Code!="200"{
		t.Error("fail testIpc resp1:" ,resp1.Code,resp1.Body," resp2:",resp2.Code,resp2.Body)
	}

	fmt.Println("test finish")

	//client1.Close()
	client2.Close()
}
*/
