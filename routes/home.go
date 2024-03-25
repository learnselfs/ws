package routes

import "github.com/learnselfs/whs"

func homeRoutes(home *whs.Route) {
	home.GET("base", homeBaseContext)
	home.GET("index", homeIndexContext)
	//home.RegisterRouter("header", homeHeaderContext)
}

func homeBaseContext(c *whs.Context) {
	c.Html(200, "base.html", map[string]interface{}{"title": "Base"})
}
func homeIndexContext(c *whs.Context) {
	c.Html(200, "index.html", map[string]interface{}{"title": "index"})
}
func homeHeaderContext(c *whs.Context) {
	c.Html(200, "header.html", map[string]interface{}{"title": "header"})
}
