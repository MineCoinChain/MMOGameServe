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
	//将玩家对象添加到世界管理器中
	core.WorldMgrObj.AddPlayer(npc)
	fmt.Println("----> player ID = ", npc.Pid, "Online...", ", Player num = ", len(core.WorldMgrObj.Players))

}
func main() {
	s := znet.NewServer("MMOGame")

	//hook函数
	s.AddOnConnStart(OnConnectionStart)
	//注册路由

	s.Serve()

}
