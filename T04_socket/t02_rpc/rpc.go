package rpc

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

//注册服务对象并开启该 RPC 服务的代码如下：
func server() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func client() {
	// 此时， RPC 服务端注册了一个Arith类型的对象及其公开方法Arith.Multiply()和
	// Arith.Divide()供 RPC 客户端调用。 RPC 在调用服务端提供的方法之前，必须先与 RPC 服务
	// 端建立连接，如下列代码所示：
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 在建立连接之后， RPC 客户端可以调用服务端提供的方法。首先，我们来看同步调用程序顺
	// 序执行的方式：
	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
	// 此外，还可以以异步方式进行调用，具体代码如下：
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done
}
