package main

import (
	"fmt"
)

type testtype struct {
	a int
	b string
}

func (t *testtype) string() string {
	return fmt.Sprint("aaa:", t.a) + " " + t.b
}
func main() {
	t := &testtype{77, "set"}
	fmt.Println(t.string())
}
