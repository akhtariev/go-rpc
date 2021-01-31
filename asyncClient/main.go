package main

import (
	"fmt"
	gorpc "go-rpc"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost" + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Asynchronous call
	quotient := new(gorpc.Quotient)
	args := &gorpc.Args{7,8}
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	<-divCall.Done	// will be equal to divCall
	// check errors, print, etc.
	if divCall.Error != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d/%d=%d", args.A, args.B, divCall.Reply)
}
