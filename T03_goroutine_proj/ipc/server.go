package ipc

import (
	"encoding/json"
	"fmt"
)

//Server
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

//IpcServer
type IpcServer struct {
	Server
}

//Request
type Request struct {
	Method string "method"
	Params string "params"
}

//response
type Response struct {
	Code string "code"
	Body string "body"
}


//NewIpcServer()
func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

//connect()
func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" { // 关闭该连接
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b) // 返回结果
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}


