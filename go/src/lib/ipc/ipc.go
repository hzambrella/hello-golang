package ipc
//ipc frame
import(
	"encoding/json"
	"fmt"
)

type Request struct{
	Method string "method"
	Params string "params"
}

type Response struct{
	Code string "code"
	Body string "body"
}

type Server interface{
	Name() string
	Handle (method,params string)*Response
}

type IpcServer struct{
	Server
}

func NewIpcServer(server Server)*IpcServer{
	return &IpcServer{server}
}

func (server *IpcServer)Connect()chan string{
		session:=make(chan string,0)
		go func (c chan string){
		for {
			fmt.Println("ipc.go:35:new for")
			request:=<-c
			if request=="CLOSE"{
				fmt.Println("ipc.go:38:break")// main procedure is finish ,it is useless
				break   // close this connect
			}
			var req Request
			var resp *Response
			err:=json.Unmarshal([]byte(request),&req)
			if err!=nil{
				// TODO:deal err
				resp=&Response{"500",err.Error()}
				panic(err)
			}

			resp=server.Handle(req.Method,req.Params)
			b,err:=json.Marshal(resp)
			c<-string(b)  //return result

			fmt.Println(req.Method,":[",resp.Code,"]")
		}
		fmt.Println("ipc.go:55:thread end")
	}(session)

	fmt.Println("ipc.go:57:A new session(channal) has been created successsfully " )
	return session
}


