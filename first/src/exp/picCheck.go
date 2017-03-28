package main
import (
	"fmt"
//	"strconv"
	"strings"
//	"math"
//	"log"
)

var picFormat="jpeg,jpg,png"
//use for money
func main(){
	html:="http://www.baidu.com/fadsfa.exe"
	html2:="http://www.baidu.com/fadsfa.jpeg"
	fmt.Println(pictureCheck(html))
	fmt.Println(pictureCheck(html2))
}

func pictureCheck(picLink string)(bool){
	strSlice:=strings.Split(picLink,".")

	lenStr:=len(strSlice)
	lastStr:=strSlice[lenStr-1]
	fmt.Println(lastStr)
	if lenStr>1{
		return strings.Contains(picFormat,lastStr)
	}

	return true
}
