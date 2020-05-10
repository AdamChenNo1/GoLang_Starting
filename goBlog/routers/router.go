package routers

import(
	"goBlog/controllers"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/",&controllers.MainController{})
}