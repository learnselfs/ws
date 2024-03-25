package routes

import (
	"github.com/learnselfs/ws/config"
	"github.com/learnselfs/ws/middleware"
)

var ()

func SetStatic(staticRoute, staticPath string, htmlPath string) {
	config.Server.Static(staticRoute, staticPath)
	//config.Server.Template(htmlPath)
	config.Server.TemplateDelim("[[", "]]", htmlPath)
}

func Routes() {
	admin := config.Server.Group("admin/")
	registerRoutes(admin)

	config.Server.UseMiddleware(middleware.ParseJwtHandler)

	menu := admin.Group("menu/")
	menuRoutes(menu)

	//api := config.Server.Group("api/")
	home := config.Server.Group("home/")
	homeRoutes(home)

}
