package main

import (
	"errors"
	gorpc "go-rpc"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Arith int

func (t *Arith) Multiply(args *gorpc.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *gorpc.Args, quo *gorpc.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	var exit = make(chan int)
	<-exit
}