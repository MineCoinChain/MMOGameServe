package core

import "fmt"

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
func (m *AOIManager) GetWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

//获取每个小格子y轴方向上的高度
func (m *AOIManager) GetHeight() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager : \n MinX:%d,MaxX:%d,cntsX:%d, minY:%d, maxY:%d,cntsY:%d, Grids inManager:\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	for _, v := range m.Grids {
		s += fmt.Sprintln(v)
	}
	return s
}
