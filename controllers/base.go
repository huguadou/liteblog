package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

//如果子controller存在NestPrepare()方法，就实现该接口
type NestPreparer interface {
	NestPreparer()
}
type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseController")
	//判断子类是否实现了NestPreparer接口,如果实现了就调用接口方法
	//将页面路径 保存到Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	app, ok := ctx.AppController.(NestPreparer)
	if ok {
		app.NestPreparer()
	}
}
