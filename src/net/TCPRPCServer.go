package main

import (
	"errors"
	"net/rpc"
	"net"
	"fmt"
)

/*

基于TCP的RPC服务器的实现


 */
// 参数结构体
type FactorialArgs1 struct {
	N int   // 函数的参数
}
//  用于计算阶乘的内部函数
func factorial1(n int) int {
	if n <= 1 {
		return 1
	} else {
		return factorial1(n - 1) * n
	}
}

type Factorial1 struct {

}

func (this *Factorial1) GetFactorial1(args *FactorialArgs1, reply *int) error {
	n := args.N
	if n < 0 || n > 12 {
		return errors.New("n必须是0到12之间的一个整数！")
	} else {
		*reply = factorial1(n)
	}
	return nil
}
func main() {
	f := new(Factorial1)
	rpc.Register(f)
	tcpAddr,_ := net.ResolveTCPAddr("tcp",":6543")
	listener,_ := net.ListenTCP("tcp",tcpAddr)
	fmt.Println("RPC【TCP】阶乘服务器已经启动！")
	for {
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}


}