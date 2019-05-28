package core

import (
	"sync"
	"fmt"
)

//定义格子类

type Grid struct {
	//格子ID
	GID       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	PlayerIDS map[int]interface{}
	PIDLock   sync.RWMutex
}

func NewGrid(id int, minx int, maxx int, miny int, maxy int) *Grid {
	return &Grid{
		GID:       id,
		MinX:      minx,
		MaxX:      maxx,
		MinY:      miny,
		MaxY:      maxy,
		PlayerIDS: make(map[int]interface{}),
	}
}

//添加一个玩家信息
func (this *Grid) AddPlayer(id int, player interface{}) {
	this.PIDLock.Lock()
	if _, ok := this.PlayerIDS[id]; !ok {
		this.PlayerIDS[id] = player
	}
	this.PIDLock.Unlock()
}

//删除一个玩家信息
func (this *Grid) DeletePlayer(id int) {
	this.PIDLock.Lock()
	delete(this.PlayerIDS, id)
	this.PIDLock.Unlock()
}

//得到当前格子内所有玩家ID
func (this *Grid) GetPlayers() (playerIDs []int) {
	this.PIDLock.RLock()
	for playerID, _ := range this.PlayerIDS {
		playerIDs = append(playerIDs, playerID)
	}
	this.PIDLock.RUnlock()
	return playerIDs
}

//调试打印格子的方法
func (this *Grid) String() string {
	return fmt.Sprintf("Grid id : %d, minX:%d, maxX:%d , minY:%d, maxY:%d, playerIDs:%v\n",
		this.GID, this.MinX, this.MaxX, this.MinY, this.MaxY, this.PlayerIDS)
}
