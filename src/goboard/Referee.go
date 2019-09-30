package goboard

import (
	"goutil"
)

type Referee struct {
	situation Situation
	sandTable Situation
}

func (referee *Referee) Watch(Situation [][]int) {
	referee.situation.Init(Situation)
}

func (referee *Referee) Draw(color int, size int) int {
	referee.sandTable = referee.situation.DeepCopy()
	ourCorps := goutil.NewSet()
	EnemyCorps := goutil.NewSet()
	ourArmy := goutil.NewSet()
	Enemies := goutil.NewSet()
	cf := new(CoordinateFactory).Init(size)
	//棋子参军
	for x, i := range referee.situation.GetData() {
		for y, j := range i {
			if j == color {
				ourArmy.Add(cf.GetCoordinate(x, y))
			} else if j == -color {
				Enemies.Add(cf.GetCoordinate(x, y))
			}
		}
	}
	//编制军团
	for _, i := range ourArmy.Values() {
		ourSoldier := i.(Coordinate)
		if !inSetSub(ourSoldier, ourCorps) {
			squad := goutil.NewSet()
			referee.sandTable.mass(ourSoldier, squad)
			ourCorps.Add(squad)
		}
	}
	for _, i := range Enemies.Values() {
		enemy := i.(Coordinate)
		if !inSetSub(enemy, EnemyCorps) {
			squad := goutil.NewSet()
			referee.sandTable.mass(enemy, squad)
			EnemyCorps.Add(squad)
		}
	}
	//提去死者

	return 0
}

func inSetSub(coo Coordinate, s *goutil.Set) bool {
	for _, sub := range s.Values() {
		if sub.(*goutil.Set).Contains(coo) {
			return true
		}
	}
	return false
}
