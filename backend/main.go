package main

import (
	"b2mak/theScoreAssignemnt/controllers"
	_ "b2mak/theScoreAssignemnt/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
