package main

import(
	_"goBlog/routers"
	"github.com/astaxie/beego"
)

func main(){
	beego.SetStaticPath("/static","static")
	beego.Run()
}