package core

import (
	"fmt"
	"log"
)

type AOIManager struct {
	MinX  int
	MaxX  int
	CntsX int
	MinY  int
	MaxY  int
	CntsY int
	Grids map[int]*Grid
}

//定义初始化方法

func NewAOIManager(minx int, maxx int, cntsx int, miny int, maxy int, cntsy int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minx,
		MaxX:  maxx,
		CntsX: cntsx,
		MinY:  miny,
		MaxY:  maxy,
		CntsY: cntsy,
		Grids: make(map[int]*Grid),
	}
	//初始化地图中的格子
	for x := 0; x < aoiMgr.CntsX; x++ {
		for y := 0; y < aoiMgr.CntsY; y++ {
			//计算格子ID
			gid := y*aoiMgr.CntsX + x
			aoiMgr.Grids[gid] = NewGrid(
				//格子ID
				gid,
				aoiMgr.MinX+x*aoiMgr.GetWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.GetWidth(),
				aoiMgr.MinY+y*aoiMgr.GetHeight(),
				aoiMgr.MaxY+(y+1)*aoiMgr.GetHeight(),
			)
		}
	}
	return aoiMgr
}

//获取每个小格子x轴方向上的宽度
func (this *AOIManager) GetWidth() int {
	return (this.MaxX - this.MinX) / this.CntsX
}

//获取每个小格子y轴方向上的高度
func (this *AOIManager) GetHeight() int {
	return (this.MaxY - this.MinY) / this.CntsY
}

func (this *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager : \n MinX:%d,MaxX:%d,cntsX:%d, minY:%d, maxY:%d,cntsY:%d, Grids inManager:\n",
		this.MinX, this.MaxX, this.CntsX, this.MinY, this.MaxY, this.CntsY)
	for _, v := range this.Grids {
		s += fmt.Sprintln(v)
	}
	return s
}

//添加一个playID到地图的格子中
func (this *AOIManager) AddPidToGrid(pId, gId int) {
	this.Grids[gId].AddPlayer(pId)
}

//删除地图格子中的一个ID
func (this *AOIManager) DeletePidToGrid(pId, gId int) {
	this.Grids[gId].DeletePlayer(pId)
}

//查询地图格子中的ID
func (this *AOIManager) GetPidToGrid(gId int) []int {
	return this.Grids[gId].GetPlayers()
}

//根据玩家所在格子的ID查找视野内的所有玩家ID
func (this *AOIManager) GetSurroundGridsByGid(gID int) (playerID []int) {
	if _, ok := this.Grids[gID]; !ok {
		log.Println("查找的ID不存在")
		return
	}

	//定义一个存储格子ID的切片
	var gridIDs []int
	gridIDs = append(gridIDs, gID)
	//获取gID所在格子的x和y坐标
	g_x := gID % this.CntsX
	g_y := gID / this.CntsY

	if g_x > 0 {
		gridIDs = append(gridIDs, gID-1)
	}

	if g_x < this.CntsX-1 {
		gridIDs = append(gridIDs, gID+1)
	}
	temp := gridIDs
	if g_y > 0 {
		for _, v := range temp {
			gridIDs = append(gridIDs, v-this.CntsX)
		}
	}
	if g_y < this.CntsY-1 {
		for _, v := range temp {
			gridIDs = append(gridIDs, v+this.CntsX)
		}
	}
	//TODO 临时验证
	fmt.Println(gridIDs)
	for _, v := range gridIDs {
		playerID = append(playerID, this.GetPidToGrid(v)...)
	}

	return

}

//根据玩家坐标确定玩家所在格子ID
func (this *AOIManager) GetGIDByPos(x float32, y float32) int {
	if x > float32(this.MaxX) || x < float32(this.MinX) || y < float32(this.MinY) || y > float32(this.MaxY) {
		return -1
	}
	pos_X := (int(x) - this.MinX) / this.GetWidth()
	pos_Y := (int(y) - this.MinY) / this.GetHeight()
	return pos_X + pos_Y*this.CntsX

}

//根据坐标获取一个九宫格内所有pID的集合
func (this *AOIManager) GetPIDsByPos(x float32, y float32) []int {
	gID := this.GetGIDByPos(x, y)
	return this.GetPidToGrid(gID)

}

//根据坐标将pID添加到九宫格内
func (this *AOIManager) AddPIDByPos(pID int, x float32, y float32) {
	gID := this.GetGIDByPos(x, y)
	this.AddPidToGrid(pID, gID)
}

//根据坐标将一个palyer从九宫格中删除
func (this *AOIManager) DeletePIDByPos(pID int, x float32, y float32) {
	gID := this.GetGIDByPos(x, y)
	this.DeletePidToGrid(pID, gID)
}
