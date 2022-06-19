package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (arith *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (arith *Arith) Divide(args *Args, quotient *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quotient.Quo = args.A / args.B
	quotient.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("resolve error: ", err.Error())
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("listen tcp error: ", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error: ", err.Error())
		}
		// rpc调用只需要处理整个连接，声明好方法即可，不用写具体方法调用逻辑
		rpc.ServeConn(conn)
	}
}
