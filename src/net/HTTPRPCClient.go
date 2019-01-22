package main

import (
	"net/rpc"
	"log"
	"fmt"
)

/*

func (this *Factorial) GetFactorial(args *FactorialArgs, reply *int) error {
}

client -> server：Factorial.GetFactorial     FactorialArgs对象
 */
type ClientFactorialArgs struct {
	N int
}

func main() {
	serverAddress := "127.0.0.1:5432"
	client,err := rpc.DialHTTP("tcp",serverAddress)

	if err != nil {
		log.Fatal("连接服务端错误",err)
	}

	args := ClientFactorialArgs{10}

	var reply int
	//  调用RPC服务端的GetFactorial函数
	err = client.Call("Factorial.GetFactorial",args,&reply)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Printf("阶乘：%d! = %d\n",args.N,reply)
	args.N = 15
	//  调用RPC服务端的GetFactorial函数
	err = client.Call("Factorial.GetFactorial",args,&reply)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Printf("阶乘：%d! = %d\n",args.N,reply)
}