package research
import(
	"fmt"
)

// interface
type Singer interface{
	Sing(song string)
}

func Singing(people,song string){
	var s Singer
	switch people{
		case "man":
			s=&Man{Name:"man usually",money:"100"}
		case "woman":
			s=&Woman{Dress:"mini skirt",money:"100"}
	}
	s.Sing(song)
}

//=========
type Man struct{
	Name string
	money string
}

func (m *Man)Getmoney()string{
	return m.money
}

func (m *Man)Sing(song string){
	fmt.Println(m.String())
}

func (m *Man)String()string{
	return m.Name+" say :caonimabi, lao zi  bu  hui  chang !  sha bi"
}

type Woman struct{
	Dress string
	money string
}

func (m *Woman)Getmoney()string{
	return m.money
}

func (m *Woman)Sing(song string){
	fmt.Println(m.String(song))
}

func (m *Woman)String(song string)string{
	return "woman usually dress "+m.Dress+" to sing "+song
}
