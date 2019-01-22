package main

import (
	"errors"
	"net/rpc"
	"net/http"
	"fmt"
)

/*

基于HTTP的RPC实现

TCP
UDP
WebSocket

RPC（Remote Procedure Call，远程过程调用），一种协议，像调用本地函数一样调用远程函数

RPC类似于WebService  WSDL  Stub（庄）

1. HTTP
2. TCP
3. JSON


1.  编写待调用的函数

  （1）参数，需要定义一个参数结构体（第1个参数）
  （2）返回值，指针类型变量返回参数（第2个参数）
  （3）真正的返回值是一个错误类型的值（error）

2.  注册该函数

3.  启动HTTP监听服务
 */

// 参数结构体
type FactorialArgs struct {
	N int   // 函数的参数
}
//  用于计算阶乘的内部函数
func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return factorial(n - 1) * n
	}
}

type Factorial struct {

}

func (this *Factorial) GetFactorial(args *FactorialArgs, reply *int) error {
	n := args.N
	if n < 0 || n > 12 {
		return errors.New("n必须是0到12之间的一个整数！")
	} else {
		*reply = factorial(n)
	}
	return nil
}
func main() {
	f := new(Factorial)
	rpc.Register(f)
	rpc.HandleHTTP()
	fmt.Println("RPC阶乘服务器已经启动！")
	err := http.ListenAndServe(":5432",nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}