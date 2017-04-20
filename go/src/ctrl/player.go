package ctrl

import (
	"fmt"
)

type Player struct {
	Name   string `json:"name"`
	Exp    string `json:"exp"`
	Level  int    `json:"level"`
	RoomId int    `json:"room"`

	mess chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 10)
	p := &Player{mess: m}

	go func(p *Player) {
		mRec := <-p.mess
		fmt.Println(mRec)
	}(p)

	return p
}
