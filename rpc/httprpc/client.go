package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing error: ", err.Error())
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("multiply error: ", err.Error())
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)
	//quotient := Quotient{}
	var quotient Quotient
	err = client.Call("Arith.Divide", args, &quotient)
	if err != nil {
		log.Fatal("divide error: ", err.Error())
	}
	fmt.Printf("%d/%d=%d...%d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}
