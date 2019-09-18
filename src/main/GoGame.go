package main

import (
	"fmt"
	"goboard"
)

func main() {
	var cf *goboard.CoordinateFactory = new(goboard.CoordinateFactory).Init(19)
	var gb = goboard.GetGoBoard(19)

	for i := 0; i < 3; i++ {
		gb.Set(cf.GetCoordinate(0, i))
		gb.Set(cf.GetCoordinate(1, i))
	}
	gb.Set(cf.GetCoordinate(18, 18))
	gb.Set(cf.GetCoordinate(0, 3))
	fmt.Print(gb.GetDataJsonStyle())
}
