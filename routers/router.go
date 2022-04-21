package routers

import (
	"BeegoDemo2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/", &controllers.MainController{})
	//注册接口  post接口
	beego.Router("/register", &controllers.RegisterController{})
	//登录
	beego.Router("/login", &controllers.LoginController{})

	//退出
	beego.Router("/exit", &controllers.ExitController{})

}
