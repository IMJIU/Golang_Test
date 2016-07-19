package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"io"
	"io/ioUtil"
)
func main(){
//	验证IP地址有效性的代码如下：
//	func net.ParseIP()
//	创建子网掩码的代码如下：
//	func IPv4Mask(a, b, c, d byte) IPMask
//	获取默认子网掩码的代码如下：
//	func (ip IP) DefaultMask() IPMask
//	根据域名查找IP的代码如下：
//	func ResolveIPAddr(net, addr string) (*IPAddr, error)
//	func LookupHost(name string) (cname string, addrs []string, err error)；
	main1()
}
/**
HTTP/1.1 200 OK
Server: Apache-Coyote/1.1
Accept-Ranges: bytes
ETag: W/"2608-1445934104282"
Last-Modified: Tue, 27 Oct 2015 08:21:44 GMT
Content-Type: text/html;charset=UTF-8
Content-Length: 2608
Date: Tue, 03 Nov 2015 09:51:24 GMT
Connection: close
*/
func main1() {
//	if len(os.Args) != 2 {
//		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
//		os.Exit(1)
//	}
//	service := os.Args[1]
	service := "localhost:9000"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
/**
HTTP/1.1 200 OK
Server: Apache-Coyote/1.1
Accept-Ranges: bytes
ETag: W/"2608-1445934104282"
Last-Modified: Tue, 27 Oct 2015 08:21:44 GMT
Content-Type: text/html;charset=UTF-8
Content-Length: 2608
Date: Tue, 03 Nov 2015 07:18:14 GMT
Connection: close
*/
func main2() {
	service := "localhost:9000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

