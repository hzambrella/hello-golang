package loghz

import(
	"fmt"
	"runtime"
	"time"
)

const(
	TIME_FORMAT="2006-01-02 15:04"
)

// about depth:please find api :package runtime
func fileLine(depth int)(string,int){
	_,file,line,ok:=runtime.Caller(depth)
	if !ok{
		file="?"
		line=0
	}
	return file,line
}

func Error(err error){
	if err==nil{
		return
	}
	now:=time.Now()

	file,line:=fileLine(2)

	fmt.Printf("[%s,ERROR]:%s:%d\n",now.Format(TIME_FORMAT),file,line)
}
