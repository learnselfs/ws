package routes

import (
	"github.com/learnselfs/whs"
	"github.com/learnselfs/ws/middleware"
)

func registerRoutes(register *whs.Route) {
	register.POST("login", LoginContext)
	register.POST("logout", LogoutContext)
}

func LoginContext(c *whs.Context) {
	middleware.JwtHandler(c)
}
func LogoutContext(c *whs.Context) {
	middleware.ParseJwtHandler(c)
}
