package goboard

import "goutil"

type SituationImpl struct {
	GoBoardImpl
}

type Situation interface {
	Init(situation [][]int) Situation
	Take(set *goutil.Set)
	take(coordinate Coordinate)
	DeepCopy() Situation
	GetData() [][]int
	mass(coordinate Coordinate, set *goutil.Set)
}

func (s *SituationImpl) Init(situation [][]int) Situation {
	if s.Data.Len() < 1 {
		s.Data.PushBack(situation)
	} else {
		s.Data.Remove(s.Data.Back())
		s.Data.PushBack(situation)
	}
	return s
}

func (s *SituationImpl) take(coordinate Coordinate) {
	s.Data.Back().Value.([][]int)[coordinate.Get()[0]][coordinate.Get()[1]] = s.GetColor(coordinate) * 4
}

func (s *SituationImpl) DeepCopy() Situation {
	var newS Situation = new(SituationImpl)
	newData := s.GetData()
	newS.Init(newData)
	return newS
}

func (s *SituationImpl) mass(coordinate Coordinate, set *goutil.Set) {
	color := s.GetColor(coordinate)
	if color != WHITE && color != BLACK && color != BLANK {
		return
	} //todo
	if !set.Contains(coordinate) {
		set.Add(coordinate)
		for _, i := range coordinate.AllAround() {
			if s.GetColor(i) == color {
				s.mass(i, set)
			}
		}
	}
}

func (s *SituationImpl) IsLiveSquad(squad goutil.Set) {
	libertySet := goutil.NewSet()
	for _, i := range squad.Values() {
		soldier := i.(Coordinate)
		for _, around := range soldier.AllAround() {
			if s.GetColor(around) == BLANK && !libertySet.Contains(around) {
				libertySet.Add(around)
			}
		}
	}
	//todo
}

// 死棋
// 1.隔两个空无援手 2.无两个活气
// 3一个活气且少于七个子的体量和空间(角) 少于8个子的体量和空间(边)
func (s *SituationImpl) IsDeath(squad goutil.Set) {

}
