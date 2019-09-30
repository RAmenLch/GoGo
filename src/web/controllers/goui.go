package controllers

import "github.com/astaxie/beego"

type GoUIController struct {
	beego.Controller
}

func (c *GoUIController) Get() {
	c.TplName = "goboard.html"
}
