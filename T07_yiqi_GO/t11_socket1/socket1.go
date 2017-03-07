package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)

	ips, err := net.LookupIP("www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ips)

	ip, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
