// Package routes @Author Bing
// @Date 2024/3/21 13:56:00
// @Desc
package routes

import (
	"bytes"
	"encoding/json"
	"github.com/learnselfs/whs"
	"github.com/learnselfs/wlog"
)

// RouteRecordRaw /** 路由参数简介 */
type RouteRecordRaw struct {
	/* 路由访问路径 */
	Path string `json:"path,omitempty"`
	/* 路由 Name (对应页面组件 Name, 可用作 KeepAlive 缓存标识 && 按钮权限筛选) */
	Name string `json:"name,omitempty"`
	/** 路由重定向地址 */
	Redirect string `json:"redirect,omitempty"`
	/** 视图文件路径 */
	Component string `json:"component,omitempty"`
	/** 路由元信息 */
	Meta Meta `json:"meta,omitempty"`
	/** 多级路由嵌套 */
	Children []*RouteRecordRaw `json:"children,omitempty"`
}

// Meta /** 路由元信息 */
type Meta struct {
	/** 菜单和面包屑对应的图标 */
	Icon string `json:"icon,omitempty"`
	/** 路由标题 (用作 document.title || 菜单的名称) */
	Title string `json:"title,omitempty"`
	/** 是否在菜单中隐藏, 需要高亮的 path (通常用作详情页高亮父级菜单) */
	ActiveMenu string `json:"activeMenu,omitempty"`
	/** 路由外链时填写的访问地址 */
	IsLink string `json:"isLink,omitempty"`
	/** 是否在菜单中隐藏 (通常列表详情页需要隐藏) */
	IsHide bool `json:"isHide,omitempty"`
	/** 菜单是否全屏 (示例：数据大屏页面) */
	IsFull bool `json:"isFull,omitempty"`
	/** 菜单是否固定在标签页中 (首页通常是固定项) */
	IsAffix bool `json:"isAffix,omitempty"`
	/** 当前路由是否缓存 */
	IsKeepAlive bool `json:"isKeepAlive,omitempty"`
}

func newMenu() *RouteRecordRaw {
	return &RouteRecordRaw{
		Children: make([]*RouteRecordRaw, 0),
	}
}

func newMenus() []*RouteRecordRaw {
	var buf bytes.Buffer
	menus := make([]*RouteRecordRaw, 0)
	for _, route := range routes {
		buf.WriteString(route)
		//var menu = newMenu()
		var menu RouteRecordRaw
		err := json.Unmarshal(buf.Bytes(), &menu)
		if err != nil {
			wlog.Errorf("list: %#v", menus)
		}
		menus = append(menus, &menu)
		buf.Reset()
	}
	return menus
}

func menuRoutes(menu *whs.Route) {
	menu.GET("list", listContent)
}

func listContent(c *whs.Context) {
	OKWithData(c, menus, "ok")
}
