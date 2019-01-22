package main

import (
	"log"
	"fmt"
	"net/rpc/jsonrpc"
)

/*

func (this *Factorial) GetFactorial(args *FactorialArgs, reply *int) error {
}

client -> server：Factorial.GetFactorial     FactorialArgs对象
 */
type ClientFactorialArgs2 struct {
	N int
}

func main() {
	serverAddress := "127.0.0.1:6543"
	client,err := jsonrpc.Dial("tcp",serverAddress)


	if err != nil {
		log.Fatal("连接服务端错误",err)
	}

	args := ClientFactorialArgs2{12}

	var reply int
	//  调用RPC服务端的GetFactorial函数
	err = client.Call("Factorial2.GetFactorial2",args,&reply)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Printf("阶乘：%d! = %d\n",args.N,reply)
	args.N = 15
	//  调用RPC服务端的GetFactorial函数
	err = client.Call("Factorial2.GetFactorial2",args,&reply)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Printf("阶乘：%d! = %d\n",args.N,reply)
}