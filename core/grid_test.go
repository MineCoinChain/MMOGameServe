package core

import (
	"fmt"
	"testing"
)

func TestGrid(t *testing.T) {

	player1 := "玩家1"
	player2 := "玩家2"


	//单元测试Grid模块
	g := NewGrid(1,1,2,10,20)

	g.AddPlayer(1, player1)
	g.AddPlayer(2, player2)

	//打出格子信息
	fmt.Println(g)
}