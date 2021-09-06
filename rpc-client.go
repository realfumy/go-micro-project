package main

import (
	"fmt"
	"net/rpc"
	// "net/rpc/jsonrpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "192.168.109.128:8899")
	if err != nil {
		fmt.Println("连接异常", err)
		return
	}
	defer conn.Close()
	var resp string
	err = conn.Call("fumy-micro-service.HelloMyfu", "付敏洋", &resp)
	if err != nil {
		fmt.Println("连接异常", err)
		return
	}
	defer conn.Close()
	fmt.Println(resp)
}
