package routers

import (
	"BeegoDemo2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//注册接口  post接口
	beego.Router("/register", &controllers.RegisterController{})
	//获取注册用户信息 Get接口--学习用的
	beego.Router("/register", &controllers.RegisterController{})

}
