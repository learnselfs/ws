// Package middleware @Author Bing
// @Date 2024/3/5 10:45:00
// @Desc
package middleware

import (
	"crypto/rand"
	"encoding/json"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/learnselfs/whs"
	"github.com/learnselfs/wlog"
	"net/http"
)

var (
	secret []byte
)

func init() {
	_, err := rand.Read(secret)
	if err != nil {
		wlog.Errorf("rand.Read:%s", err)
	}
}

func ParseJwtHandler(c *whs.Context) {
	// 1. 校验 jwt
	authorization := c.Request.Header.Get("X-Access-Token")
	if len(authorization) == 0 {
		return
	}
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(authorization[7:], claims, func(t *jwt.Token) (any, error) { return secret, nil })
	if err != nil {
		// 2.1 jwt 失败
		wlog.Infof("%#v", err)
		return
	}
	// 2.2 jwt 成功
	wlog.Infof("%#v", claims)
	c.Next()

}

func JwtHandler(c *whs.Context) {
	// 1. 获取用户名、密码
	// 2. 校验用户名密码
	// 3. 生成 jwt
	username := c.Request.Form.Get("username")
	password := c.Request.Form.Get("password")
	var account Account
	err := json.NewDecoder(c.Request.Body).Decode(&account)
	if err != nil {
		wlog.Errorf("error: %s", err)
		return
	}
	if account.Username == "admin" && account.Password == "admin" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{username: account.Username, password: account.Password})
		jwtString, err := token.SignedString(secret)
		if err != nil {
			wlog.Errorf("error: %s", err)
			return
		}
		data := map[string]interface{}{"Authorization": jwtString}
		c.Json(http.StatusOK, data, "ok")
		return
	}
	whs.NotFoundHandler(c)
	return
}

type Msg struct {
	Code    int            `json:"code"`
	Data    map[string]any `json:"data"`
	Message string         `json:"msg"`
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
