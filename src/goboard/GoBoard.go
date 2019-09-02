package goboard

import (
	"container/list"
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
	"goutil"
)

type GoBoardImpl struct {
	Data  *list.List
	color int // -1 is black ,1 is white
	size  int
}

type GoBoard interface {
	Init(size int)
	Set(coordinate Coordinate)
	GoBack()
	GetData() [][]int
	GetColor(coordinate Coordinate) int
	Take(set *hashset.Set)
}

func (gb *GoBoardImpl) Init(size int) {
	gb.size = size
	var value [][]int = goutil.MakeIntMatrix([2]int{size, size})
	gb.Data = list.New()
	gb.Data.PushBack(value)
	gb.color = -1
}

func (gb *GoBoardImpl) Set(coordinate Coordinate) {
	if gb.GetColor(coordinate) != 0 {
		panic(errors.New("已有子!"))
	}
	var value = (gb.Data.Back()).Value.([][]int)
	var newValue = goutil.CopyMatrix(value)
	newValue[coordinate.Get()[0]][coordinate.Get()[1]] = gb.color
	gb.Data.PushBack(newValue)
	var flag bool = false
	for _, around := range coordinate.AllAround() {
		if gb.GetColor(around) == -gb.color {
			var checked = hashset.New()
			if gb.CheckLiberty(around, -gb.color, checked) == 0 {
				gb.Take(checked)
				flag = true
			}
		} else if gb.GetColor(around) == 0 {
			flag = true
		}
	}
	if !flag {
		if gb.CheckLiberty(coordinate, gb.color, hashset.New()) <= 1 {
			panic(errors.New("自杀着!"))
		}
	}

	for prev := gb.Data.Back().Prev(); prev != nil; prev = prev.Prev() {
		if goutil.Equal(gb.Data.Back().Value.([][]int), prev.Value.([][]int)) {
			gb.GoBack()
			panic(errors.New("全局同形"))
		}
	}
	gb.color = -gb.color
}

func (gb *GoBoardImpl) GoBack() {
	gb.Data.Remove(gb.Data.Back())
	gb.color = -gb.color
}

func (gb *GoBoardImpl) GetData() [][]int {
	return gb.Data.Back().Value.([][]int)
}

func (gb *GoBoardImpl) GetColor(coordinate Coordinate) int {
	if coordinate.IsOutOfBoard() {
		return -2
	}
	return gb.Data.Back().Value.([][]int)[coordinate.Get()[0]][coordinate.Get()[1]]
}

func (gb *GoBoardImpl) CheckLiberty(coordinate Coordinate, color int, checked *hashset.Set) int {
	var libertySet = hashset.New()
	return gb.checkLiberty(coordinate, color, libertySet, checked)
}

func (gb *GoBoardImpl) checkLiberty(coordinate Coordinate, color int, libertySet *hashset.Set, checked *hashset.Set) int {
	var allAround = coordinate.AllAround()
	if checked.Contains(coordinate) {
		return libertySet.Size()
	}
	checked.Add(coordinate)
	for _, around := range allAround {
		if gb.GetColor(around) == 0 {
			libertySet.Add(around)
		} else if gb.GetColor(around) == color {
			gb.checkLiberty(around, color, libertySet, checked)
		}
	}
	return libertySet.Size()
}
func (gb *GoBoardImpl) Take(set *hashset.Set) {
	var values = set.Values()
	for _, value := range values {
		gb.take(value.(Coordinate))
	}
}

func (gb *GoBoardImpl) take(coordinate Coordinate) {
	gb.Data.Back().Value.([][]int)[coordinate.Get()[0]][coordinate.Get()[1]] = 0
}

func GetGoBoard(size int) GoBoard {
	var gb GoBoard = new(GoBoardImpl)
	gb.Init(19)
	return gb
}
