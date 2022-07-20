package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err!=nil{
		log.Fatal("dail:",err)
	}
	args:=Args{7,8}
	var reply int
	err =client.Call("Arith.Mul",args,&reply)
	if err!=nil{
		log.Fatal("call",err)
	}
	fmt.Printf("Arith:%d*%d=%d",args.A,args.B,reply)
}
