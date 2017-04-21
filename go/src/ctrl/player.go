package ctrl

import (
	"fmt"
)

type Player struct {
	Name   string  `json:"name"`
	Exp    float64 `json:"exp"`
	Level  int     `json:"level"`
	RoomId int     `json:"room"`

	mess chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 10)
	p := &Player{mess: m}

	go func(p *Player) {
		fmt.Println("player.go:21:new thread")
		for {
			fmt.Println("player.go:23:new for")
			mRec := <-p.mess
			if mRec.To == "all" {
				if mRec.From == "admin" {
					fmt.Println(fmt.Sprintf("[all]system:[%s]", mRec.Content))
				} else {
					fmt.Println(fmt.Sprintf("[all]%s:[%s]", mRec.From, mRec.Content))
				}
			}

			if mRec.To == p.Name {
				fmt.Println(fmt.Sprintf("[private]%s:%s", mRec.From, mRec.Content))
			}
			fmt.Println("player.go:35:for end")

			if mRec.From==p.Name&&mRec.To==p.Name{
				if mRec.Content=="///system:quit"{
					break
				}
			}
		}//for end

		fmt.Println("player.go:33: thread end")
	}(p)

	return p
}
