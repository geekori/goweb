package main

import (
    "os"
    "fmt"
)

/*

文件和目录操作


 */

func main() {
	// 创建目录

    //os.Mkdir("mydir",0777)
    //os.MkdirAll("mydir/test/hello/abc/123",0777)
    //  待删除的目录必须为空
   /* err := os.Remove("mydir")
    if err != nil {
        fmt.Println(err)
    }
    // 可以删除不为空的目录
    err = os.RemoveAll("mydir")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("mydir成功删除")
    }*/

    //  文件操作

    //  文件写操作
    userfile := "hello.txt"
    /*fout, err := os.Create(userfile)
    if err != nil {
        fmt.Println(userfile,err)
        return
    }
    defer fout.Close()

    for i := 0; i < 20; i++ {
        fout.WriteString("hello world\r\n")
        fout.Write([]byte("How are you?\r\n"))
    }*/

    //  读文件
    fl,err := os.Open("hello.txt")

    if err != nil {
        fmt.Println(userfile,err)
        return
    }
    defer fl.Close()
    buf := make([]byte,1024)
    for {
        n,_:=fl.Read(buf)
        if n == 0 {
            break
        }
        os.Stdout.Write(buf[:n])
    }

    // 删除文件
    os.Remove("hello.txt")
}
