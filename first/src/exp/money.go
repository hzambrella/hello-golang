package main
import (
	"fmt"
	"strconv"
	"strings"
//	"math"
	"log"
)

//use for money
func main(){
//	price:=19.10932
//	price:=10000000019.1
	price:=float64(10019)
	fmt.Println(price)
	str:=strconv.FormatFloat(price,'f',-2,64)
//	str1:=strconv.ParseFloat(price,64)
	fmt.Println(str)
//	fmt.Println(string(price))
	pricenow,err:=strconv.ParseFloat(str,64)
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println(pricenow)
	strSlice:=strings.Split(str,".")
	fmt.Println(strSlice)
	for _,v:=range strSlice{
		fmt.Println(v)
	}
	if len(strSlice)>1{
	if len(strSlice[1])>1{
		fmt.Println("warning!")
	}
}

}
