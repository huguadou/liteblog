package routers

import (
	"liteblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//  beego.Router("/", &controllers.MainController{})
	//注解路由 需要调用Include
	beego.Include(&controllers.IndexController{})
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
}
