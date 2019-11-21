package controllers

import (
	"beegoApi/models"
	"log"

	"github.com/astaxie/beego"
)

//如果子controller存在NestPrepare()方法，就实现该接口
type NestPreparer interface {
	NestPreparer()
}
type BaseController struct {
	beego.Controller
	IsLogin bool        //标识 用户是否登录
	User    models.User //登录的用户
}

//定义session中的key
const SESSION_USER_KEY = "SESSION_USER_KEY"

func (ctx *BaseController) Prepare() {
	log.Println("BaseController")
	//判断子类是否实现了NestPreparer接口,如果实现了就调用接口方法
	//将页面路径 保存到Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	app, ok := ctx.AppController.(NestPreparer)
	if ok {
		app.NestPreparer()
	}
	//验证用户是否登录，判断session中是否存在用户，存在就已经登录，不存在就没有登录
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY)
	if tu != nil {
		u, ok := tu.(models.User)
		if ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
}
