package controllers

import (
	"errors"
	"liteblog/models"
	"strings"
)

type UserController struct {
	BaseController
}

//@router /login [post]
func (c *UserController) Login() {
	/*c.GetMustString(arg0,arg1,string) 是在BaseController里面定义的，第一个参数获取请求的的参数键值对的key,请求后，如key对于的value是空，就返回第二个参数*/
	//判断邮箱不能空
	email := c.GetMustString("email", "邮箱不能为空！")
	//判断密码不能为空
	pwd := c.GetMustString("password", "密码不能为空！")
	var (
		user *models.User
		err  error
	)
	user, err := models.QueryUserByEmailAndPassword(email, pwd)
	if err != nil {
		c.Abort500(syserrors.NewError("邮箱或密码不对", err))
	}
	//将user保存到session
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONOk("登录成功", "/")
}

//@router /reg [poset]
func (c *UserController) Reg() {
	name := c.GetMustString("name", "昵称不能为空！")
	email := c.GetMustString("email", "邮箱不能为空!")
	pwd1 := c.GetMustString("password", "密码不能为空！")
	pwd2 := c.GetMustString("password2", "确认密码不能为空！")
	if strings.Compare(pwd1, pwd2) != 0 {
		c.Abort500(errors.New("密码与确认密码 必须要一致!"))
	}
	u, err := models.QueryUserByName(name)
	if err == nil && u.ID != 0 {
		c.Abort500(syserrors.NewError("用户昵称已经存在",err))
	}
	u,err:=models.QueryUserByEmail(email)
	if err ==nil && u.ID !=0{
		c.Abort500(syserrors.NewError("用户邮箱已经存在！"，err))
	}
	//开始保存用户
	err:=models.SaveUser(
		&models.User{
			Name: name,
			Email:email,
			Pwd:pwd1,
			Avatar:"",
			Role:1,
		}
	)
	if err !=nil{
		c.Abort500(sysnerror.NewError("用户注册失败",err))
	}
	c.JSONOk("注册成功"，"/user")
}
