package cg

import (
	"encoding/json"
	"errors"
	"ipc"
)

// CenterClient
type CenterClient struct {
	*ipc.IpcClient
}

//增加一个用户
func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}
	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

//移除一个用户
func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removeplayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

//遍历用户列表
func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

//组播
func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message} // 构造Message结构体
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}