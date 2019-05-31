package core

import (
	"sync"
)

//当前世界地图的边界参数
const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)



//当前场景的世界管理模块
type WorldManager struct {
	//当前全部在线的Player集合
	Players map[int32]*Player
	//保护player集合的锁
	pLock sync.RWMutex
	//AOIManager当前的地图管理器件
	AoiMgr *AOIManager
}
//对外提供一个全局世界管理模块的指针
var WorldMgrObj *WorldManager

func init() {
	//创建一个全局的世界管理对象
	WorldMgrObj = NewWorldManager()
}

//初始化方法
func NewWorldManager() *WorldManager {
	return &WorldManager{
		Players: make(map[int32]*Player),
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
	}

}

//添加一个玩家
func (this *WorldManager) AddPlayer(player *Player) {
	//添加到在线用户集合中，因为是并发的map操作，所以必须枷锁

	this.pLock.Lock()

	this.Players[player.Pid] = player

	this.pLock.Unlock()

	//添加到世界地图中
	this.AoiMgr.AddPIDByPos(int(player.Pid), player.X, player.Z)

}

//删除一个玩家
func (this *WorldManager) RemovePlayerByID(pID int32) {
	this.pLock.Lock()
	//先从地图中删除

	player := this.GetPlayerByID(pID)
	this.AoiMgr.DeletePIDByPos(int(pID), player.X, player.Z)
	//在从当前在线用户集合中删除
	delete(this.Players, pID)
	this.pLock.Unlock()

}

//通过一个玩家ID 得到一个Player对象
func (this *WorldManager) GetPlayerByID(pID int32) *Player {
	this.pLock.RLock()
	defer this.pLock.RUnlock()

	if p, ok := this.Players[pID]; ok {
		return p
	}


	return nil
}

//获取全部的在线玩家集合
func (this *WorldManager) GetAllPlayers() [] *Player {
	this.pLock.RLock()

	var onLinePlayer []*Player
	//遍历当前在线用户集合并返回
	for _, v := range this.Players {
		onLinePlayer = append(onLinePlayer, v)
	}

	this.pLock.RUnlock()

	return onLinePlayer
}
