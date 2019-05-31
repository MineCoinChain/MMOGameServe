package apis

import (
	"zinx/znet"
	"zinx/ziface"
	"MMOGameServe/pb"
	"github.com/golang/protobuf/proto"
	"log"
	"MMOGameServe/core"
)

type MoveRouter struct {
	znet.BaseRouter
}

func (m *MoveRouter) Handler(request ziface.IRequest) {
	var pos = &pb.Position{}
	if err := proto.Unmarshal(request.GetMessage().GetMsgData(), pos); err != nil {
		log.Println("proto err:", err)
		return
	}
	//获取发送位置信息的玩家
	 pid, err := request.GetConnection().GetProperty("PID")
	 if err != nil {
		log.Println("moverouter request err:",err)
		return
	}
	player:=core.WorldMgrObj.GetPlayerByID(pid.(int32))
	//将消息发送给其余所有玩家
	player.UpdatePosition(pos.X,pos.Y,pos.Z,pos.V)

}
