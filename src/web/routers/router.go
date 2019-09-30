package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	beego.Router("/", &controllers.GoUIController{})
	beego.Router("/go", &controllers.GoBoardController{})
}
