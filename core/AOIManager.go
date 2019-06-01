package core

import (
	"fmt"
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

//输出函数的方法
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
//通过一个格子ID得到当前格子的周边九宫格的格子ID集合
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gid是否在AOI中
	if _, ok := m.Grids[gID]; !ok {
		return
	}

	//将当前中心GID放入九宫格切片中
	grids = append(grids, m.Grids[gID])

	//==== > 判读GID左边是否有格子？ 右边是否有格子

	//通过格子ID 得到x轴编号 idx = gID % cntsX
	idx := gID % m.CntsX

	//判断idx编号左边是否还有格子
	if idx > 0 {
		//将左边的格子加入到 grids 切片中
		grids = append(grids, m.Grids[gID-1])
	}

	//判断idx编号右边是否还有格子
	if idx < m.CntsX-1 {
		//将右边的格子加入到 grids 切片中
		grids = append(grids, m.Grids[gID+1])
	}

	// ===> 得到一个x轴的格子集合，遍历这个格子集合
	// for ... 依次判断  格子的上面是否有格子？下面是否有格子

	//将X轴全部的Grid ID 放到一个slice中 ，遍历整个slice
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	for _, gid := range gidsX {
		//10,11,12
		//通过Gid得到当前Gid的Y轴编号
		//idy = gID / cntsX
		idy := gid/m.CntsX

		//上方是否还有格子
		if idy > 0 {
			grids = append(grids, m.Grids[gid-m.CntsX])
		}
		//下方是否还有格子
		if idy < m.CntsY-1 {
			grids = append(grids, m.Grids[gid+m.CntsX])
		}
	}

	return
}

//通过x，y坐标得到对应的格子ID
func (m *AOIManager) GetGidByPos(x, y float32) int {
	if x < 0 || int(x) >= m.MaxX {
		return -1
	}
	if y < 0 || int(y) >= m.MaxY {
		return -1
	}
	//根据坐标 得到当前玩家所在格子ID
	idx := (int(x) - m.MinX) / m.GetWidth()
	idy := (int(y) - m.MinY) / m.GetHeight()

	//gid  = idy*cntsX + idx
	gid := idy * m.CntsX + idx


	return gid
}

//根据一个坐标 得到 周边九宫格之内的全部的 玩家ID集合
func (m *AOIManager) GetSurroundPIDsByPos(x, y float32) (playerIDs []int) {

	//通过x，y得到一个格子对应的ID
	gid := m.GetGIDByPos(x,y)

	//通过格子ID 得到周边九宫格 集合
	grids := m.GetSurroundGridsByGid(gid)

	fmt.Println("gid = ", gid)

	//将分别将九宫格内的全部的玩家 放在 playerIDs
	for  _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayers()...)
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
