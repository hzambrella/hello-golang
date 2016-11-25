package main

import (
	"fmt"
)

type node struct {
	id  int
	pid int
	p   *node
	c1  int
	c2  int
	c3  int
}

func (x *node) bind(y *node) {

	x.c1++

	// x.p != y：避免二环 1-2-1
	if x.p != y {
		x.c2 += y.c1
	}

	// 避免二环和 x.p.p != y 三环 1-2-3-1
	if x.p == nil || (x.p != nil && x.p != y && x.p.p != y) {
		x.c3 += y.c2
	}

	// 避免二环1-2-1
	if x.p != nil && x.p != y {
		x.p.c2 += 1
		// 避免三环
		if x.p.p != y {
			x.p.c3 += y.c1
		}
		// 避免三环
		if x.p.p != nil && x.p.p != y {
			x.p.p.c3 += 1
		}
	}
	// 将绑定点的parent设置为x.
	y.p = x
	y.pid=x.id

}

// 打印
func (x *node) print() {
	fmt.Printf("%d:  pid:%d  1级: %d  2级:%d   3级: %d\n", x.id, x.pid, x.c1, x.c2, x.c3)
}

func main() {
	hztest()
	// from sql
	fmt.Println("模拟 sql")
	x1 := &node{id: 1, pid: 0, c1: 1, c2: 1, c3: 1}
	sql(x1)
	x1 = &node{id: 1, pid: 0, c1: 1, c2: 1, c3: 0}
	sql2(x1)
}

func sql(xIn *node) {
	// from sql
	x2 := &node{id: 2, pid: 1, c1: 1, c2: 1, c3: 0}
	x3 := &node{id: 3, pid: 2, c1: 1, c2: 0, c3: 0}
	x4 := &node{id: 4, pid: 3, c1: 1, c2: 0, c3: 0}
	// 预处理
	nodebind(x2, x3)
	nodebind(x3, x4)
	x4.bind(xIn)
xIn.print()
	x2.print()
	x3.print()
	x4.print()
	fmt.Println("--------")
}

func sql2(xIn *node) {
	// from sql
	x1 := &node{id: 1, pid: 0, c1: 1, c2: 1, c3: 0}
	x2 := &node{id: 2, pid: 1, c1: 1, c2: 0, c3: 0}
	x3 := &node{id: 3, pid: 2, c1: 0, c2: 0, c3: 0}
	// 预处理
	nodebind(x1, x2)
	nodebind(x2, x3)
	x3.bind(xIn)
	x1.print()
xIn.print()
	x2.print()
	x3.print()
	fmt.Println("--------")
}

func nodebind(x1 *node, x2 *node) {
	if x2.pid == x1.id {
		x2.p = x1
	} else {
		fmt.Println("not parent")
	}
}
func hztest() {
	test1()
	test2()
	test3()
	// from sql
	x1 := &node{id: 1, pid: 0, c1: 0, c2: 0, c3: 0}
	sql(x1)

}

func test1() {
	x1 := &node{id: 1, pid: 0, c1: 0, c2: 0, c3: 0}
	x2 := &node{id: 2, pid: 0, c1: 0, c2: 0, c3: 0}
	x3 := &node{id: 3, pid: 0, c1: 0, c2: 0, c3: 0}
	x4 := &node{id: 4, pid: 0, c1: 0, c2: 0, c3: 0}
	x1.bind(x2)
	x2.bind(x3)
	x3.bind(x4)
	x4.bind(x1)
	x1.print()
	x2.print()
	x3.print()
	x4.print()
	fmt.Println("--------")
}

func test2() {
	x1 := &node{id: 1, pid: 0, c1: 0, c2: 0, c3: 0}
	x2 := &node{id: 2, pid: 0, c1: 0, c2: 0, c3: 0}
	x3 := &node{id: 3, pid: 0, c1: 0, c2: 0, c3: 0}
	x1.bind(x2)
	x2.bind(x3)
	x3.bind(x1)
	x1.print()
	x2.print()
	x3.print()
	fmt.Println("--------")
}

func test3() {
	node1 := &node{id: 1, pid: 0, c1: 0, c2: 0, c3: 0}
	node2 := &node{id: 2, pid: 0, c1: 0, c2: 0, c3: 0}
	node3 := &node{id: 3, pid: 0, c1: 0, c2: 0, c3: 0}
	node4 := &node{id: 4, pid: 0, c1: 0, c2: 0, c3: 0}
	node5 := &node{id: 5, pid: 0, c1: 0, c2: 0, c3: 0}
	node6 := &node{id: 6, pid: 0, c1: 0, c2: 0, c3: 0}
	node7 := &node{id: 7, pid: 0, c1: 0, c2: 0, c3: 0}
	node8 := &node{id: 8, pid: 0, c1: 0, c2: 0, c3: 0}
	node1.bind(node2)
	node3.bind(node4)
	node4.bind(node6)
	node6.bind(node5)
	node5.bind(node7)
	node7.bind(node1)
	node2.bind(node3)
	node2.bind(node8)
	node1.print()
	node2.print()
	node3.print()
	node4.print()
	node5.print()
	node6.print()
	node7.print()
	node8.print()
	fmt.Println("-----------")
}

