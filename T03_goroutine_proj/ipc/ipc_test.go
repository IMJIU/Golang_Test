package ipc

import (
	"testing"
	"fmt"
)

type EchoServer struct {}

func (server *EchoServer) Handle(method, params string) *Response  {
	fmt.Println("ECHO:" + method + ",param:"+params)
	return &Response{Body:"ECHO:" + method ,Code: "200"}
}
func (server *EchoServer) Name() string {
	return "EchoServer"
}
func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1,err1 := client1.Call("broadcast","From Client1")
	if err1 != nil {
		fmt.Println("Failed Call", err1)
	}
	resp2,err2 := client1.Call("broadcast","From Client2")
	if err2 != nil {
		fmt.Println("Failed adding player", err2)
	}
	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1.Body, "resp2:", resp2.Body)
	}
	client1.Close()
	client2.Close()
}
