package proto

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	// 私人
	Default = iota
	// 公屏
	Public = iota
	//服务，设置姓名
	SetName = iota
)

const (
	ServerUid = "0"
	PublicUid = "-1000"
)

var (
	RecieveError = errors.New("wrong reciever")
	NoNameError  = errors.New("noname")
	NoUidError   = errors.New("nouid")
)

type User struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

type Message struct {
	UidFrom    string `json:"from"`
	UidTo      string `json:"to"`
	Content    string `json:"message"`
	Type       int    `json:"type"`
	CreateTime string `json:"create_time"`
}

func NewUser(uid string, name string) (*User, error) {
	if len(uid) == 0 {
		return nil, NoUidError
	}
	u := &User{uid, name}
	return u, nil
}

func (u *User) MakeMess(messtype int, uidTo, content string) ([]byte, error) {
	t1 := time.Now().Format("2006-01-02")
	if messtype == Public {
		uidTo = PublicUid
	}

	mess := &Message{
		UidFrom:    u.Uid,
		UidTo:      uidTo,
		Content:    content,
		Type:       messtype,
		CreateTime: t1,
	}
	b, err := json.Marshal(mess)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (u *User) GetMess(b []byte) (*Message, error) {
	m := Message{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("json", string(b))
		return nil, err
	}

	/*
		if m.UidFrom == "0" || m.UidTo == "0" {
			//do nothing
		} else {
			if m.UidTo != PublicUid && m.UidTo != u.Uid {
				return nil, RecieveError
			}
		}
	*/
	return &m, nil
}

func (u *User) String() string {
	return fmt.Sprintf("(%s)%s\n", u.Uid, u.Name)
}

func (m *Message) String() string {
	switch m.Type {
	case Default:
		return fmt.Sprintf("[from %s](%v):%s\n ", m.UidFrom, m.CreateTime, m.Content)
	case Public:
		return fmt.Sprintf("[PUBLIC][from %s](%v):%s\n ", m.UidFrom, m.CreateTime, m.Content)
	case SetName:
		return fmt.Sprintf("[SETNAME][from %s](%v):%s\n ", m.UidFrom, m.CreateTime, m.Content)

	default:
		return "wrong message type!"
	}
}
