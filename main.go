package main

//在routers包前面，可以看到有一个“_”，这表明是引入routers包，并执行init方法
import (
	_ "BeegoDemo2/routers"
	"BeegoDemo2/util"

	"github.com/astaxie/beego"
)

func main() {
	util.InitMysql()
	beego.Run()
}
