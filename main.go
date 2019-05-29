package main

import (
	"zinx/znet"
	"zinx/ziface"
	"fmt"
	"MMOGameServe/core"
)

//定义一个hook函数
func OnConnectionStart(conn ziface.IConnection) {
	fmt.Println("测试前后台链接是否正常")
	npc := core.NewPlayer(conn)
	npc.ReturnPid()
	npc.ReturnPlayerPosition()


}
func main() {
	s := znet.NewServer("MMOGame")

	//hook函数
	s.AddOnConnStart(OnConnectionStart)
	//注册路由

	s.Serve()

}
