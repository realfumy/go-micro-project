package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Myfu struct {
}

func (myfu *Myfu) HelloMyfu(name string, resp *string) error {
	*resp = "你好" + name
	return nil
}
func main() {
	err := rpc.RegisterName("fumy-micro-service", new(Myfu))
	if err != nil {
		fmt.Println("注册服务失败", err)
		return
	}
	listener, err := net.Listen("tcp", "192.168.109.128:8899")
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("监听异常")
		return
	}
	defer conn.Close()
	fmt.Println("监听成功！")
	rpc.ServeConn(conn)
}
