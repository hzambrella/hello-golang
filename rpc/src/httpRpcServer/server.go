package httpRpcServer
import(
	"net/http"
	"net/rpc"
	"fmt"
)

type Args struct{
	A,B int
}

type Reply struct{
	Result int
}

type MathRpc int//for rpc  register


func (m *MathRpc)Mutiply(args Args,reply *Reply)error{
	reply.Result=args.A*args.B
	return nil
}

func ServerRpc(){
	mathRpc:=new(MathRpc)
	rpc.Register(mathRpc)
	rpc.HandleHTTP()

	fmt.Println("at :1234")
	err:=http.ListenAndServe(":1234",nil)
	if err!=nil{
		fmt.Println(err.Error())
	}
}
