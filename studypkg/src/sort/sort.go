package main

import (
	"sort"
	"fmt"
	"strconv"
)

type A []string

func main() {
	var a A=[]string{"10","3","6","8","9","2"}
	sort.Sort(a)
	fmt.Println(a)
}

func (a A)Less(i,j int)bool{
	x,err:=strconv.Atoi(a[i])
	if err!=nil{
		panic(err)
	}
	y,err:=strconv.Atoi(a[j])
	if err!=nil{
		panic(err)
	}

	//return x<y//if x>y 
	return x>y
}

func (a A)Len()int{
	return len(a)
}

func (a A)Swap(i,j int){
	a[i],a[j]=a[j],a[i]
}
