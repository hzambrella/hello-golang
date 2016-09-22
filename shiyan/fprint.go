// 研究Sprint
package main

import "fmt"

const (
	d = `d=%s`
)

func main() {
	a := "1"
	b := "2"
	c := fmt.Sprint("a=" + a + "b=" + b)
	fmt.Println(c)
	e := fmt.Sprintf(d, a)
	fmt.Println(e)
}
