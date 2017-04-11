package ctrl
import(
	"model"
)
//Cnc[i] 表示第i号顶点在第几个连通分量当中
type Cnc map[int64]int

func SubGraph()(error) {
	// 当前连通分量
	cnt:=0
	cnc:=make(Cnc)

	dAll:=model.New()
	allnum:=len(dAll)
}
