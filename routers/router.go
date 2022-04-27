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
	//显示文章详情
	beego.Router("/article/:id", &controllers.ShowArticleController{})

	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})

	//上传文件
	beego.Router("/upload", &controllers.UploadController{})

	//更新文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})

	// 删除文章
	beego.Router("/article/delete", &controllers.DeleteArticleController{})

	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController2{})

	//标签
	beego.Router("/tags", &controllers.TagsController{})

}
