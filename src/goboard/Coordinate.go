package goboard

import "errors"
import "strconv"

type CoordinateImpl struct {
	x    int
	y    int
	size int
}

type Coordinate interface {
	Get() []int
	Set(x int, y int)
	setSize(size int)
	North() Coordinate
	South() Coordinate
	West() Coordinate
	East() Coordinate
	IsOutOfBoard() bool
	AllAround() []Coordinate
}

func (coordinate *CoordinateImpl) Set(x int, y int) {
	if x >= coordinate.size || y >= coordinate.size || x < 0 || y < 0 {
		panic(errors.New("超出棋盘大小范围Coordinate=(" + strconv.Itoa(x) + "," + strconv.Itoa(y) + "),size=" + strconv.Itoa(coordinate.size)))
	}
	coordinate.x = x
	coordinate.y = y
}

func (coordinate *CoordinateImpl) Get() []int {
	return []int{coordinate.x, coordinate.y}
}
func (coordinate *CoordinateImpl) setSize(size int) {
	coordinate.size = size
}

func (coordinate *CoordinateImpl) North() Coordinate {
	var north = new(CoordinateImpl)
	north.x = coordinate.x - 1
	north.y = coordinate.y
	north.size = coordinate.size
	if north.x < 0 {
		north.x = -1
		north.y = -1
	}
	return north
}

func (coordinate *CoordinateImpl) South() Coordinate {
	var south = new(CoordinateImpl)
	south.x = coordinate.x + 1
	south.y = coordinate.y
	south.size = coordinate.size
	if south.x >= coordinate.size {
		south.x = -1
		south.y = -1
	}
	return south
}

func (coordinate *CoordinateImpl) East() Coordinate {
	var east = new(CoordinateImpl)
	east.x = coordinate.x
	east.y = coordinate.y + 1
	east.size = coordinate.size
	if east.y >= coordinate.size {
		east.x = -1
		east.y = -1
	}
	return east
}

func (coordinate *CoordinateImpl) West() Coordinate {
	var west = new(CoordinateImpl)
	west.x = coordinate.x
	west.y = coordinate.y - 1
	west.size = coordinate.size
	if west.y < 0 {
		west.x = -1
		west.y = -1
	}
	return west
}

func (coordinate *CoordinateImpl) AllAround() []Coordinate {
	var x []Coordinate = make([]Coordinate, 4)
	x[0] = coordinate.North()
	x[1] = coordinate.East()
	x[2] = coordinate.South()
	x[3] = coordinate.West()
	return x
}

func (coordinate *CoordinateImpl) IsOutOfBoard() bool {
	return coordinate.x < 0 // x ==-1,y==-1
}

type CoordinateFactory struct {
	size int
}

func (cf *CoordinateFactory) Init(size int) *CoordinateFactory {
	cf.size = size
	return cf
}

func (cf *CoordinateFactory) GetCoordinate(x int, y int) Coordinate {
	var coo Coordinate = new(CoordinateImpl)
	coo.setSize(cf.size)
	coo.Set(x, y)
	return coo
}
