package ipc
import(
	"encoding/json"
)

const(
	ServerErr="500"
)

type IpcClient struct{
	conn chan string
}

func NewIpcClient(server *IpcServer)*IpcClient{
	c:=server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient)Call(method,params string)(resp *Response,err error){
	req:=&Request{method,params}
	var b []byte
	b,err=json.Marshal(req)
	if err!=nil{
		//return &Response(ServerErr,err.Error()),err
		return
	}

	client.conn<-string(b)
	str:=<-client.conn// wait return value

	var resp1 Response
	if err=json.Unmarshal([]byte(str),&resp1);err!=nil{
		//return &Response(ServerErr,err.Error()),err
		return
	}

	resp=&resp1
	//return resp,nil
	return
}

func(client *IpcClient)Close(){
	client.conn<-"CLOSE"
}
