package apis

import (
	"zinx/znet"
	"zinx/ziface"
	"github.com/golang/protobuf/proto"
	"MMOGameServe/pb"
	"log"
	"MMOGameServe/core"
)

type WorldChat struct {
	znet.BaseRouter
}

func (this *WorldChat) Handler(request ziface.IRequest) {
	//获取客户端发送的数据
	var content = &pb.Talk{}
	if err := proto.Unmarshal(request.GetMessage().GetMsgData(), content); err != nil {
		log.Println("proto unmarshal eror :", err)
		return
	}
	//将数据传递给各个客户端

	pid, err := request.GetConnection().GetProperty("PID")
	if err != nil {
		log.Println("获取PID错误:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByID(pid.(int32))
	//将消息发送给所有玩家
	player.SendMsgToAll(content.GetContent())

}
