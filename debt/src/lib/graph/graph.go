// based on  http://github.com/tleham/graph.git
// I add something about:
//1:depth first search(dfs)
//2:can use dfs to find connected components 
package graph

import (
	"bytes"
	"fmt"
	"strconv"
	"errors"
	"github.com/tlehman/ds"
)
var ErrEdgeNotFound=errors.New("edge not found")

// forms a linked list of incident edges
type edge struct {
	y      int
	Weight float64
	next   *edge
}

// Each vertex is identified by the index into the edges array
//    AdjList
//      0 -> { y: 1 } -> { y: 3}
//      1 -> { y: 2 } -> { y: 3}
//      2 -> { y: 0 }

type AdjList struct {
	edges       []*edge
	edgeCount   int
	vertexCount int
	directed    bool
	weight		bool
	vertexset   []bool
}

// Builds a new adjaceny list, storing the boolean flag that
// distinguishes a directed from an undirected graph.
func New(directed,weight bool) AdjList {
	edges := make([]*edge, 0)
	return AdjList{edges: edges, directed: directed,weight:weight}
}

// Returns the size of the set of vertices
// Runs in O(1) time
func (g AdjList) VertexCount() int {
	return g.vertexCount
}

// Returns the size of the set of edges
// Runs in O(1) time
func (g AdjList) EdgeCount() int {
	return g.edgeCount
}


// Add edge adds an edge node to the x-th spot in the edges slice.
// It resizes if necessary.
func (g *AdjList) addEdge(x, y int) {
	if max(x,y) >= len(g.edges) {
		g.resizeEdges(max(x,y))
	}
	if g.edges[x].y == -1 {
		g.edges[x].y = y
		g.edges[x].next = nil
	} else {
		newnext := g.edges[x]
		g.edges[x] = &edge{y: y, next: newnext}
	}
	g.edgeCount += 1
	// record the vertices
	if !g.vertexset[x] {
		g.vertexset[x] = true
		g.vertexCount += 1
	}
	if !g.vertexset[y] {
		g.vertexset[y] = true
		g.vertexCount += 1
	}
}

func (g *AdjList) AddEdge(x, y int) {
	g.addEdge(x,y)
	if !g.directed {
		g.addEdge(y,x)
	}
}

func (g *AdjList) AddEdgeWithWeight(x, y int,weight float64) {
	g.weight=true
	g.addEdge(x,y)
	g.AddWeight(x,y,weight)
	if !g.directed {
		g.addEdge(y,x)
		g.AddWeight(y,x,weight)
	}
}

func (g *AdjList)AddWeight(x,y int,weight float64)error{
	g.weight=true
	notFound:=true
	for e:=g.edges[x];e!=nil;e=e.next{
		if e.y==y&&e.Weight==0{
			e.Weight=weight
			notFound=false
		}
	}

	if notFound{
		errMess:=fmt.Sprintf("edge <%d %d> not found",x,y)
		return errors.New(errMess)
	}
	return nil
}
/*
func (g *AdjList)GetWeight(x,y int)(float64,error){
	notFound:=true
	var weight float64
	for e:=g.edges[x];e!=nil;e=e.next{
		if e.y==y{
			weight=e.weight
			notFound=false
		}
	}

	if notFound{
		errMess:=fmt.Sprintf("edge <%d %d> not found",x,y)
		return 0,errors.New(errMess)
	}
	return weight,nil
}
*/
func (g *AdjList) resizeEdges(size int) {
	diff := size - len(g.edges)
	for i := 0; i <= diff; i++ {
		g.edges = append(g.edges, &edge{y: -1})
		g.vertexset = append(g.vertexset, false)
	}
}

// Builds a graphviz string representing the graph 
func (g AdjList) String() string {
	var buffer bytes.Buffer
	var arrow string
	if g.directed {
	    buffer.WriteString("di")
		arrow = "->"
	} else {
		// the strict keyword combines multiple edges in the rendering
		// undirected graphs are modeled using both edges for convenience
	    buffer.WriteString("strict ")
		arrow = "--"
	}
	buffer.WriteString("graph {\n")
	for x, e := range g.edges {
		if x > -1 && e.y > -1 {
			for c := e; c != nil; c = c.next {
				if g.weight{
					weightStr:=strconv.FormatFloat(c.Weight,'f',-1,64)
					buffer.WriteString(fmt.Sprintf("  %d %s %d[label=%s];\n", x, arrow, c.y,weightStr))
				}else{
					buffer.WriteString(fmt.Sprintf("  %d %s %d;\n", x, arrow, c.y))
				}
			}
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Find connected components in an undirected graph.
// The return value of Components() is a slice of integers
// that associates each vertex with a component number.
// if fs==true use bfs ,else use dfs
func (g *AdjList) Components(fs bool) []int {
	// each index v is a vertex, and the value comps[v]
	// is the number of the component v is in

	comps := make([]int, len(g.edges))
	discovered := make([]bool, len(g.edges))
	compNum := 0
	idcsComp:=make([]int ,0)

	if fs{
		fmt.Println("use Breadth-First Search for undirected graph")
	}else{
		fmt.Println("use Depth-First Search for undirected graph")
	}

	for v, _ := range g.edges {
		if !discovered[v] {
			// indices of component containing v
			if fs{
				idcsComp = g.bfs(v)
			}else{
				idcsComp=g.dfs(v)
			}
			for _, x := range idcsComp {
				comps[x] = compNum
				discovered[x] = true
			}
			compNum++
		}
	}

	return comps
}



const (
	undiscovered = iota
	discovered = iota
	processed = iota
)

// do a Breadth-First Search from v, return indices in same
// component as v
func (g *AdjList) bfs(v int) []int {
	state := make([]int, len(g.edges)) // all initially undiscovered
	indices := make([]int, 0)
	q := ds.NewQueue()
	q.Enqueue(v)
	for q.Len() > 0 {
		x := q.Dequeue().(int)
		indices = append(indices, x)
		for e := g.edges[x]; e != nil; e = e.next {
			if e.y > -1 && state[e.y] == undiscovered {
				state[x] = discovered
				q.Enqueue(e.y)
			}
		}
		state[x] = processed
	}
	return indices
}

// do a  Depth-First Search from v, return indices in same
// component as v
func (g *AdjList) dfs(v int) []int {
	state := make([]int, len(g.edges)) // all initially undiscovered
	indices := make([]int, 0)
	state,indices=dfsfor(v,g,state,indices)
	return indices
}

func dfsfor(v int,g *AdjList,state,indices []int)([]int,[]int){
	indices=append(indices,v)
	state[v]=discovered
	// test if you need
	/*
	for _,v:=range g.edges{
		fmt.Println(v)
	}
	fmt.Println(v,state,indices)
	*/
	//
	for e:=g.edges[v];e!=nil;e=e.next{
		//test if you need
		//fmt.Printf("v %d,change %d\n ",v,e.y)
		//
		if e!=nil&&e.y>=0{
			//fmt.Println(e)
			if state[e.y]==undiscovered {
				state,indices=dfsfor(e.y,g,state,indices)
			}else{
			}
		}
	}
	return state,indices
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func MaxIntSliceElement(i []int)int{
	var m int
	for _,v:=range i{
		m=max(v,m)
	}
	return m
}
