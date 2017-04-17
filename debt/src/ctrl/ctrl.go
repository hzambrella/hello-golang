package ctrl

import(
	//"log"
	"fmt"

	"lib/graph"
	"engine/data"
)

var vrgfile="vrg"
var pngfile="png"
var significant=2
type GAL []*graph.AdjList

//get connective subgraph
func SubGraph(){
	var mainGraph="main"
	var subGraph="subgraph"
	datas:=data.New()
	directed:=true
	g := graph.New(directed)

	for _,v:=range datas.DStore{
		g.AddEdge(int(v.OwnerId),int(v.DebtorId))
	}

	if err:=draw(mainGraph,g.String());err!=nil{
		panic(err)
	}

	comps:=g.Components(true)
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
			adj:=graph.New(directed)
			gal[key]=&adj
		}
		gal[key].AddEdge(int(v.OwnerId),int(v.DebtorId))
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
