package main

import "flag"
import "fmt"
import "net/rpc"

func main() {
	addrVar := flag.String("addr", "127.0.0.1:1234", "address")
	flag.Parse()
	client, err := rpc.Dial("tcp", *addrVar)
	if err != nil {
		panic(err)
	}
	var reply uint64
	err = client.Call("Flaked.Next", uint64(0), &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%b\n", reply)
}
