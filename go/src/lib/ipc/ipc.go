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
			request:=<-c
			if request=="CLOSE"{
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

			fmt.Println("session closed.")
		}
	}(session)

	fmt.Println("A new session has been created successsfully " )
	return session
}


