package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"goboard"
)

type GoBoardController struct {
	beego.Controller
}

var gb goboard.GoBoard

func (c *GoBoardController) Get() {
	gb = goboard.GetGoBoard(19)
	c.Data["data"] = gb.GetDataJsonStyle()
	c.TplName = "api.html"
	fmt.Print(gb.GetDataJsonStyle())
}

func (c *GoBoardController) Post() {
	x, _ := c.GetInt("x")
	y, _ := c.GetInt("y")
	var cd = new(goboard.CoordinateFactory).Init(19).GetCoordinate(x, y)
	defer func() {
		err := recover()
		c.Data["data"] = gb.GetDataJsonStyle()
		c.Data["err"] = err
		c.TplName = "api.html"
	}()
	gb.Set(cd)
	c.Data["data"] = gb.GetDataJsonStyle()
	fmt.Print(gb.GetDataJsonStyle())
	c.TplName = "api.html"
}
