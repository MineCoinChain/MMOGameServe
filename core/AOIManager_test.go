package core

import (
	"testing"
)

func TestAOIManager(t *testing.T) {
	//fmt.Println(NewAOIManager(0, 250, 5, 0, 250, 5))

}

func TestGetSurroundGridsByGid(t *testing.T) {
	//初始化Manager
	n:=NewAOIManager(0, 250, 5, 0, 250, 5)
	n.GetSurroundGridsByGid(n.GetGIDByPos(2,3))
}
