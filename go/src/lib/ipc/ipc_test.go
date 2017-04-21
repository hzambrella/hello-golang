package ipc
import(
	"testing"
	"fmt"
)
var _ Server=&myServer{}

type myServer struct{
	Params string
}

func (m *myServer)Name()string{
	return "haozhao"
}

func (m *myServer)Handle(method,params string)*Response{
	return &Response{method,params}
}

func TestMyServer(t *testing.T){
	myS:=&myServer{"hahaha"}
	ipcS:=NewIpcServer(myS)
	ipcC:=NewIpcClient(ipcS)
	resp,err:=ipcC.Call("ha","ja")
	if err!=nil{
		t.Fatal()
	}
	fmt.Println(resp)
}
