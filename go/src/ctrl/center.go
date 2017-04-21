package ctrl

import (
	"encoding/json"
	"errors"
	"fmt"
	"lib/ipc"
	"sync"
)

var _ ipc.Server = &CenterServer{} // make sure that achieve Server interface

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type CenterServer struct {
	players []*Player
	//rooms   []*Room
	mutex sync.RWMutex
}

func NewCenterServer() *CenterServer {
	players := make([]*Player, 0)
	return &CenterServer{players: players}
}

func (server *CenterServer) addPlayer(params string) error {
	p := NewPlayer()
	err := json.Unmarshal([]byte(params), &p)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, p)
	return nil
}

func (server *CenterServer) removePlayer(params string) error {
	var p string
	err := json.Unmarshal([]byte(params), &p)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	for k, v := range server.players {
		if v.Name == p {
			v.mess<-&Message{From:v.Name,To:v.Name,Content:"///system:quit"}
			server.players = append(server.players[:k], server.players[k+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("player %s not found", p))
}

func (server *CenterServer) listPlayer() (string, error) {
	if len(server.players) == 0 {
		return "", errors.New("no player online, go home to eat shit")
	}
	b, err := json.Marshal(server.players)
	return string(b), err
}

func (server *CenterServer) broadcast(params string) error {
	var mess *Message
	if err := json.Unmarshal([]byte(params), &mess); err != nil {
		return err
	}

	if len(server.players) == 0 {
		return errors.New("no player online, go home to eat shit")
	} else {
		for _, v := range server.players {
			v.mess <- mess
		}
	}
	return nil
}

func (server *CenterServer) Handle(method, params string) *ipc.Response {
	resp := &ipc.Response{}
	switch method {
	case "addplayer":
		if err := server.addPlayer(params); err != nil {
			resp.Code = "500"
			resp.Body = err.Error()
		} else {
			resp.Code = "200"
		}
	case "removeplayer":
		if err := server.removePlayer(params); err != nil {
			resp.Code = "500"
			resp.Body = err.Error()
		} else {
			resp.Code = "200"
		}
	case "listplayer":
		str, err := server.listPlayer()
		if err != nil {
			resp.Code = "500"
			resp.Body = err.Error()
		} else {
			resp.Code = "200"
			resp.Body = str
		}
	case "broadcast":
		if err := server.broadcast(params); err != nil {
			resp.Code = "500"
			resp.Body = err.Error()
		} else {
			resp.Code = "200"
		}
	default:
		resp.Code = "404"
		resp.Body = "invalid request"
	}
	return resp
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}

