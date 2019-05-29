package main

import "zinx/znet"

func main() {
	s := znet.NewServer("MMOGame")

	//hook函数

	//注册路由

	s.Serve()

}
