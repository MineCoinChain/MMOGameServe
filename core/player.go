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
	Pid  int32              //玩家ID
	Conn ziface.IConnection //当前玩家的链接(与对应客户端通信)
	X    float32            //平面的x轴坐标
	Y    float32            //高度
	Z    float32            //平面的y轴坐标
	V    float32            //玩家脸朝向的方向
}

// playerID 生成器
var PidGen int32 = 1  //用于生产玩家ID计数器
var IdLock sync.Mutex //保护PidGen生成器的互斥锁

//初始化玩家的方法
func NewPlayer(conn ziface.IConnection) *Player {
	//分配一个玩家ID
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	//创建一个玩家对象
	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), //随机生成玩家上线所在的x轴坐标
		Y:    100,
		Z:    float32(140 + rand.Intn(10)), //随机在140坐标点附近 y轴坐标上线
		V:    0,                            //角度为0
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
		fmt.Println("Player send error!", err)
		return err
	}

	return nil
}

//服务器给客户端发送玩家初始ID
func (p *Player) ReturnPid() {
	//定义个msg:ID 1  proto数据结构
	proto_msg := &pb.SyncPid{
		Pid: p.Pid,
	}

	//将这个消息 发送给客户端
	p.SendMsg(1, proto_msg)
}

//服务器给客户端发送一个玩家的初始化位置信息
func (p *Player) ReturnPlayerPosition() {
	//组建MsgID:200消息
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //2 -坐标信息
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	//将这个消息 发送给客户端
	p.SendMsg(200, proto_msg)
}

//服务器给所有玩家发送消息
func (p *Player) SendMsgToAll(content string) {
	protoMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1, //Tp: 1 世界聊天, 2 坐标, 3 动作, 4 移动之后坐标信息更新
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	//遍历所有玩家集合
	players := WorldMgrObj.GetAllPlayers()
	for _, player := range players {
		player.SendMsg(200, protoMsg)
	}
}

//将自己的消息同步给周边的玩家
func (p *Player) GetSurroundingPlayers() []*Player {
	pids := WorldMgrObj.AoiMgr.GetSurroundPIDsByPos(p.X, p.Z)
	fmt.Println("Surrounding players = ", pids)
	var players []*Player
	for _, pid := range pids {
		player := WorldMgrObj.GetPlayerByID(int32(pid)) //锁粒度控制的问题
		players = append(players, player)
	}
	return players
}

//服务器向周边玩家发送消息
func (p *Player) SyncSurrounding() {

	players := p.GetSurroundingPlayers()
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}

	var proto_players []*pb.Player
	//将周边玩家的信息同步给当前玩家
	for _, player := range players {
		proto_player := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		proto_players = append(proto_players, proto_player)

	}
	proto_SyncPlayers := &pb.SyncPlayers{
		Ps: proto_players,
	}
	p.SendMsg(202, proto_SyncPlayers)

}

//格子切换的时候更新视野
func (p *Player) OnExchangeAoiGrid(oldGrid int, newGrid int) {
	//获取旧的九宫格的成员
	oldGrids := WorldMgrObj.AoiMgr.GetSurroundGridsByGid(oldGrid)
	var oldGridsMap = make(map[int]*Grid)
	for _, grid := range oldGrids {
		oldGridsMap[grid.GID] = grid
	}
	//获取新的九宫格的成员
	newGrids := WorldMgrObj.AoiMgr.GetSurroundGridsByGid(newGrid)
	var newGridMap = make(map[int]*Grid)
	for _, grid := range newGrids {
		newGridMap[grid.GID] = grid
	}

	//处理视野消失
	proto_leavemsg := &pb.SyncPid{
		Pid: p.Pid,
	}
	//在旧的视野中出现但没有在新的视野中出现的全部玩家
	var leaveingGids []int
	for gid, _ := range oldGridsMap {
		if _, ok := newGridMap[gid]; !ok {
			leaveingGids = append(leaveingGids, gid)
		}
	}
	for _, gid := range leaveingGids {
		pids := WorldMgrObj.AoiMgr.GetPidToGrid(gid)
		for _, pid := range pids {
			//在其他玩家视野中消失
			player := WorldMgrObj.GetPlayerByID(int32(pid))
			player.SendMsg(201, proto_leavemsg)
			//在自身视野中剔除旧的格子上的玩家
			proto__anoleavemsg := &pb.SyncPid{
				Pid: int32(pid),
			}
			p.SendMsg(201, proto__anoleavemsg)

		}
	}

	//处理视野出现
	proto_OnlineMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//在新的视野中出现的，但是在旧的九宫格中没有的玩家
	var newAppearGird []int
	for gid, _ := range newGridMap {
		if _, ok := oldGridsMap[gid]; !ok {
			newAppearGird = append(newAppearGird, gid)
		}
	}

	for _, gid := range newAppearGird {
		pids:=WorldMgrObj.AoiMgr.GetPidToGrid(gid)

		for _,pid:=range pids{
			//让自己出现在其他人视野中
			player:=WorldMgrObj.GetPlayerByID(int32(pid))
			player.SendMsg(200, proto_OnlineMsg)

			//让其他人出现在自己的视野中
			another_online_msg := &pb.BroadCast{
				Pid:player.Pid,
				Tp:2,
				Data:&pb.BroadCast_P{
					&pb.Position{
						X:player.X,
						Y:player.Y,
						Z:player.Z,
						V:player.V,
					},
				},
			}

			p.SendMsg(200, another_online_msg)
		}

	}

}

//更新位置信息
func (p *Player) UpdatePosition(X, Y, Z, V float32) {
	oldGrid := WorldMgrObj.AoiMgr.GetGIDByPos(p.X, p.Z)
	newGrid := WorldMgrObj.AoiMgr.GetGIDByPos(X, Z)
	if oldGrid != newGrid {
		//将旧的PID从网格中删除
		WorldMgrObj.AoiMgr.DeletePIDByPos(int(p.Pid), p.X, p.Z)
		//将新的PID添加到网格中
		WorldMgrObj.AoiMgr.AddPidToGrid(int(p.Pid), newGrid)
		//视野消失的业务
		p.OnExchangeAoiGrid(oldGrid, newGrid)
	}
	//获取周围玩家集合
	players := p.GetSurroundingPlayers()
	p.X = X
	p.Y = Y
	p.Z = Z
	p.V = V

	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}

}

//玩家下线通知
func (p *Player) OffLine() {
	proto_msg := &pb.SyncPid{
		Pid: p.Pid,
	}
	players := p.GetSurroundingPlayers()
	for _, player := range players {
		player.SendMsg(201, proto_msg)
	}
	//将玩家从地图抹除
	WorldMgrObj.RemovePlayerByID(p.Pid)
	//将玩家从世界管理器消除
	WorldMgrObj.AoiMgr.DeletePIDByPos(int(p.Pid), p.X, p.Z)

}
