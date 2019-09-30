package goboard

import (
	"container/list"
	"encoding/json"
	"errors"
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
	GetDataJsonStyle() string
	GetColor(coordinate Coordinate) int
	Take(set *goutil.Set)
}

func (gb *GoBoardImpl) Init(size int) {
	gb.size = size
	var value [][]int = goutil.MakeIntMatrix([2]int{size, size})
	gb.Data = list.New()
	gb.Data.PushBack(value)
	gb.color = BLACK
}

func (gb *GoBoardImpl) Set(coordinate Coordinate) {
	if gb.GetColor(coordinate) != BLANK {
		panic(errors.New("已有子!"))
	}
	var value = (gb.Data.Back()).Value.([][]int)
	var newValue = goutil.CopyMatrix(value)
	newValue[coordinate.Get()[0]][coordinate.Get()[1]] = gb.color
	gb.Data.PushBack(newValue)
	gb.color = -gb.color
	//验证提子,气
	var taked bool = false
	for _, around := range coordinate.AllAround() {
		z := gb.GetColor(around)
		if z == gb.color {
			liberty, checked := gb.CheckSelfLiberty(around)
			if liberty == 0 {
				gb.Take(checked)
				taked = true
			}
		}
	}
	if !taked {
		a, _ := gb.CheckSelfLiberty(coordinate)
		if a == 0 {
			gb.GoBack()
			panic(errors.New("自杀着"))
		}
	}

	//验证劫,全局同形
	for prev := gb.Data.Back().Prev(); prev != nil; prev = prev.Prev() {
		if goutil.Equal(gb.Data.Back().Value.([][]int), prev.Value.([][]int)) {
			gb.GoBack()
			panic(errors.New("全局同形"))
		}
	}

}

func (gb *GoBoardImpl) GoBack() {
	if gb.Data.Len() <= 1 {
		return
	}
	gb.Data.Remove(gb.Data.Back())
	gb.color = -gb.color
}

func (gb *GoBoardImpl) GetData() [][]int {
	//深拷贝
	return goutil.CopyMatrix(gb.Data.Back().Value.([][]int))
}

func (gb *GoBoardImpl) GetDataJsonStyle() string {
	data, _ := json.Marshal(gb.GetData())
	return string(data)
}

func (gb *GoBoardImpl) GetColor(coordinate Coordinate) int {
	if coordinate.IsOutOfBoard() {
		return EDGE
	}
	return gb.Data.Back().Value.([][]int)[coordinate.Get()[0]][coordinate.Get()[1]]
}

func (gb *GoBoardImpl) CheckSelfLiberty(coordinate Coordinate) (int, *goutil.Set) {
	if gb.GetColor(coordinate) == BLANK || gb.GetColor(coordinate) == EDGE {
		return -1, nil
	}

	check := goutil.NewSet()
	libertySet := goutil.NewSet()
	gb.mass(coordinate, check)
	for _, i := range check.Values() {
		for _, around := range i.(Coordinate).AllAround() {
			if gb.GetColor(around) == BLANK {
				libertySet.Add(around)
			}
		}
	}
	return libertySet.Size(), check
}

func (gb *GoBoardImpl) mass(coordinate Coordinate, set *goutil.Set) {
	color := gb.GetColor(coordinate)
	if color != WHITE && color != BLACK {
		return
	}
	if !set.Contains(coordinate) {
		set.Add(coordinate)
		for _, i := range coordinate.AllAround() {
			if gb.GetColor(i) == color {
				gb.mass(i, set)
			}
		}
	}
}

func (gb *GoBoardImpl) Take(set *goutil.Set) {
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
