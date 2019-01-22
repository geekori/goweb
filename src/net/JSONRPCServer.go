package main

import (
	"errors"
	"net/rpc"
	"net"
	"fmt"
	"net/rpc/jsonrpc"
)

/*

基于TCP的RPC服务器的实现


 */
// 参数结构体
type FactorialArgs2 struct {
	N int   // 函数的参数
}
//  用于计算阶乘的内部函数
func factorial2(n int) int {
	if n <= 1 {
		return 1
	} else {
		return factorial2(n - 1) * n
	}
}

type Factorial2 struct {

}

func (this *Factorial2) GetFactorial2(args *FactorialArgs2, reply *int) error {
	n := args.N
	if n < 0 || n > 12 {
		return errors.New("n必须是0到12之间的一个整数！")
	} else {
		*reply = factorial2(n)
	}
	return nil
}
func main() {
	f := new(Factorial2)
	rpc.Register(f)
	tcpAddr,_ := net.ResolveTCPAddr("tcp",":6543")
	listener,_ := net.ListenTCP("tcp",tcpAddr)
	fmt.Println("RPC【JSON】阶乘服务器已经启动！")
	for {
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}


}