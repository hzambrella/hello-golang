package loghz

import(
	"fmt"
	"runtime"
	"time"
	"strings"
)

const(
	TIME_FORMAT="2006-01-02 15:04"
)

//var IfPrintln bool=false

// about depth:please find api :package runtime
//output func name ,file name,line num
func fileLine(depth int)(string,string,int){
	pc,file,line,ok:=runtime.Caller(depth)
	var funcname string=""
	if !ok{
		file="?"
		line=0
	}else{
		funcname=runtime.FuncForPC(pc).Name()
	}
	return funcname,file,line
}

// println，显示行号和列号，方便调试
func Println(t ...interface{}){
	/*
	if !IfPrintln{
		return
	}
	*/

	funcname,file,line:=fileLine(2)

	slash:=strings.LastIndex(file,"/")
	if slash>=0{
		file=file[slash+1:]
	}

	slash2:=strings.LastIndex(funcname,"/")
	if slash2>=0{
		funcname=funcname[slash2+1:]
	}
	fmt.Printf("[[PRINT]%s:%d:'%s()']",file,line,funcname)
	fmt.Println("--->",t)

}
// 错误输出。带行号，方便检查错误和异常
func Error(err error){
	if err==nil{
		return
	}
	now:=time.Now()

	_,file,line:=fileLine(2)

	fmt.Printf("[%s,ERROR]:%s:%d\n",now.Format(TIME_FORMAT),file,line)
}
