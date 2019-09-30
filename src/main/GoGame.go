package main

import (
	"goboard"
)

func main() {
	var cf *goboard.CoordinateFactory = new(goboard.CoordinateFactory).Init(19)
	var gb = goboard.GetGoBoard(19)

	//for i := 0; i < 2; i++ {
	//	gb.Set(cf.GetCoordinate(0, i))
	//	gb.Set(cf.GetCoordinate(1, i))
	//}
	//gb.Set(cf.GetCoordinate(18, 18))
	//gb.Set(cf.GetCoordinate(0, 2))
	gb.Set(cf.GetCoordinate(0, 0))
	gb.Set(cf.GetCoordinate(0, 1))
	gb.Set(cf.GetCoordinate(1, 0))
	gb.Set(cf.GetCoordinate(1, 1))
	gb.Set(cf.GetCoordinate(18, 18))
	gb.Set(cf.GetCoordinate(2, 0))
}
