package main

//在routers包前面，可以看到有一个“_”，这表明是引入routers包，并执行init方法
import (
	_ "BeegoDemo2/routers"

	"github.com/astaxie/beego"
	//测试数据库连接
	_ "BeegoDemo2/models"
)

func main() {
	beego.Run()
}
