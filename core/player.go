
/*
	玩家模块
 */
package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
	"zinx/ziface"
	"MMOGameServe/pb"
)

type Player struct {
	Pid int32 //玩家ID
	Conn ziface.IConnection //当前玩家的链接(与对应客户端通信)
	X float32 //平面的x轴坐标
	Y float32 //高度
	Z float32 //平面的y轴坐标
	V float32 //玩家脸朝向的方向
}

// playerID 生成器
var PidGen int32 = 1//用于生产玩家ID计数器
var IdLock sync.Mutex //保护PidGen生成器的互斥锁

//初始化玩家的方法
func NewPlayer(conn ziface.IConnection) *Player{
	//分配一个玩家ID
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	//创建一个玩家对象
	p := &Player{
		Pid:id,
		Conn:conn,
		X: float32(160 + rand.Intn(10)),//随机生成玩家上线所在的x轴坐标
		Y:0,
		Z: float32(140 + rand.Intn(10)), //随机在140坐标点附近 y轴坐标上线
		V:0,//角度为0
	}

	return p
}


//玩家可以和对端客户端发送消息的方法
func (p *Player) SendMsg(msgID uint32, proto_struct proto.Message) error {
	//要将proto结构体 转换成 二进制的数据
	binary_proto_data, err := proto.Marshal(proto_struct)
	if err != nil {
		fmt.Println("marshal proto struct error ", err)
		return err
	}


	//再调用zinx原生的connecton.Send（msgID, 二进制数据）
	if err := p.Conn.Send(msgID, binary_proto_data); err != nil {
		fmt.Println("Player send error!" , err)
		return err
	}

	return nil
}

/*
 服务器给客户端发送玩家初始ID
 */
func (p *Player) ReturnPid() {
	//定义个msg:ID 1  proto数据结构
	proto_msg := &pb.SyncPid{
		Pid:p.Pid,
	}

	//将这个消息 发送给客户端
	p.SendMsg(1, proto_msg)
}

//服务器给客户端发送一个玩家的初始化位置信息
func (p *Player) ReturnPlayerPosition() {
	//组建MsgID:200消息
	proto_msg := &pb.BroadCast{
		Pid:p.Pid,
		Tp:2, //2 -坐标信息
		Data: &pb.BroadCast_P{
			P:&pb.Position{
				X:p.X,
				Y:p.Y,
				Z:p.Z,
				V:p.V,
			},
		},
	}

	//将这个消息 发送给客户端
	p.SendMsg(200, proto_msg)
}


