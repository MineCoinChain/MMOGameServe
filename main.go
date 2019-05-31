package main

import (
	"zinx/znet"
	"zinx/ziface"
	"fmt"
	"MMOGameServe/core"
	"MMOGameServe/apis"
	"log"
)

//定义一个hook函数
func OnConnectionStart(conn ziface.IConnection) {
	fmt.Println("测试前后台链接是否正常")
	npc := core.NewPlayer(conn)
	npc.ReturnPid()
	npc.ReturnPlayerPosition()
	//将玩家对象添加到世界管理器中
	core.WorldMgrObj.AddPlayer(npc)
	//给链接设置属性
	conn.SetProperty("PID", npc.Pid)
	//添加周围玩家信息
	npc.SyncSurrounding()
	fmt.Println("----> player ID = ", npc.Pid, "Online...", ", Player num = ", len(core.WorldMgrObj.Players))

}

//玩家下线后广播
func OnConnectionStop(conn ziface.IConnection) {
	pid, err := conn.GetProperty("PID")
	if err != nil {
		log.Println("conn stop err:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByID(pid.(int32))
	player.OffLine()
}
func main() {
	s := znet.NewServer("MMOGame")
	//hook函数
	s.AddOnConnStart(OnConnectionStart)
	s.AddOnConnStop(OnConnectionStop)
	//注册路由
	s.AddRouter(2, &apis.WorldChat{})
	s.AddRouter(3, &apis.MoveRouter{})
	s.Serve()

}
