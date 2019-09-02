package main

import (
	"fmt"
	"goboard"
)

func main() {
	var cf *goboard.CoordinateFactory = new(goboard.CoordinateFactory).Init(19)
	var c = cf.GetCoordinate(3, 3)
	var gb = goboard.GetGoBoard(19)
	gb.Set(c)
	gb.Set(cf.GetCoordinate(0, 0))
	gb.Set(cf.GetCoordinate(0, 1))
	gb.Set(cf.GetCoordinate(18, 18))
	gb.Set(cf.GetCoordinate(1, 0))
	fmt.Print(gb.GetData())

}
