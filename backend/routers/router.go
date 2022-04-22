package routers

import (
	"b2mak/theScoreAssignemnt/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/teams", &controllers.MainController{}, "get:GetTeams")
	beego.Router("/download", &controllers.MainController{}, "get:GetFile")
}
