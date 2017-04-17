package model
import(
	"engine/data"
)

type Debt struct{
	Id int64
	OwnerId int64	// zhai quan ren
	DebtorId int64	// zhai wu ren
	Amount float64  // amount
}

type DebtList []*Debt

func New()*Debt{
	datas:=data.New()
	for k,v:=range datas.DStore{
	}
}
