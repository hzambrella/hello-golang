package ctrl

import(
	//"log"
	"fmt"

	"lib/graph"
	"engine/data"
)

var vrgfile="vrg"
var pngfile="png"
var directed=true
var	weight=true
var significant=2
type GAL []*graph.AdjList

//get connective subgraph
func SubGraph(){
	var mainGraph="main"
	var subGraph="subgraph"
	datas:=data.New()
	g := graph.New(directed,weight)

	for _,v:=range datas.DStore{
		g.AddEdgeWithWeight(int(v.OwnerId),int(v.DebtorId),v.Amount)
	}

	if err:=draw(mainGraph,g.String());err!=nil{
		panic(err)
	}

	comps:=g.Components(false)
	gal:=subGraphAdjList(datas,g,comps,directed)
	for k,v:=range gal{
		if v!=nil{
			if v.EdgeCount()>=significant{
				fname:=fmt.Sprintf("%s_%d",subGraph,k)
				draw(fname,v.String())
			}
		}
	}
}

func subGraphAdjList(datas *data.DataS,g graph.AdjList,comps []int,directed bool)GAL{
	m:=graph.MaxIntSliceElement(comps)
	gal:=make(GAL,m+1)
	for _,v:=range datas.DStore {
		key:=comps[v.OwnerId]
		if gal[key]==nil{
			adj:=graph.New(directed,weight)
			gal[key]=&adj
		}
		gal[key].AddEdgeWithWeight(int(v.OwnerId),int(v.DebtorId),v.Amount)
	}
	return gal
}

func draw(fname,content string)error{
	_,err:=graph.SaveToFileVrg(vrgfile,fname,content)
	if err!=nil{
		return err
	}

	_,err=graph.DrawVrg(vrgfile,pngfile,fname)
	if err!=nil{
		return err
	}
	return nil
}
