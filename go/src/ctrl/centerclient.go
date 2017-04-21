package ctrl

import (
	"encoding/json"
	"errors"
	"lib/ipc"
)

type CenterClient struct {
	//wrong:ipc.IpcClient
	*ipc.IpcClient
}

//wrong,anonymous combination
/*
func NewCenterClient(server *IpcServer)*CenterClient{
	is:=ipc.NewIpcClient(server)
	return &CenterClient{is}
}
*/

func (client *CenterClient) AddPlayer(player *Player) error {
	params, err := json.Marshal(player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addplayer", string(params)) //because of anonymous
	if err != nil {
		return err
	}

	if resp.Code != "200" {
		return errors.New(resp.Body)
	}

	return nil
}

func (client *CenterClient) RemovePlayer(name string) error {
	params, err := json.Marshal(name)
	if err != nil {
		return err
	}

	resp, err := client.Call("removeplayer", string(params)) //because of anonymous
	if err != nil {
		return err
	}

	if resp.Code != "200" {
		return errors.New(resp.Body)
	}

	return nil
}

func (client *CenterClient) ListPlayer() ([]*Player, error) {

	resp, err := client.Call("listplayer", "") //because of anonymous
	if err != nil {
		return nil, err
	}
	if resp.Code != "200" {
		return nil, errors.New(resp.Body)
	}

	var players []*Player

	if err := json.Unmarshal([]byte(resp.Body), &players); err != nil {
		return nil, err
	}

	return players, nil
}

func (client *CenterClient) Broadcast(message string) error {
	mess := &Message{}
	mess.Content = message
	mess.From = "admin"
	mess.To = "all"
	params, err := json.Marshal(mess)
	if err != nil {
		return err
	}

	resp, err := client.Call("broadcast", string(params)) //because of anonymous
	if err != nil {
		return err
	}

	if resp.Code != "200" {
		return errors.New(resp.Body)
	}

	return nil
}

