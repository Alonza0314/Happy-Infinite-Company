package handlers

import (
	"encoding/json"
	"hic/models"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 登入狀態下進入登入頁面需要redirect到subpage
func LoginRedirect(c *gin.Context) {
	if c.Request.Method == "GET" && IsLogin(c) && c.Request.URL.Path == "/login/" {
		c.Redirect(http.StatusFound, "/subpage")
		c.Abort()
		return
	}
	c.Next()
}

// 需要登入的頁面在未登入狀態下需要redirect到login頁面
func NoneLoginRedirect(c *gin.Context) {
	if c.Request.Method == "GET" && !IsLogin(c) && !CheckPagePathDoNotNeedLogin(c) {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	c.Next()
}

// 確認頁面是否需要登入
func CheckPagePathDoNotNeedLogin(c *gin.Context) bool {
	paths := []string{
		"/",
		"/about/",
		"/about/history",
		"/about/members",
		"/nothing",
		"/contact",
		"/login/",
		"/login/findpw",
		"/login/resetpw",
		"/api/",
		"/api/captcha",
	}
	for _, path := range paths {
		if c.Request.URL.Path == path {
			return true
		}
	}
	return false
}

// 確認是否登入了: 是否有設定userid的cookie，確認記錄時間是否超過5分鐘
func IsLogin(c *gin.Context) bool {
	cookie, err := c.Request.Cookie("userid")
	if err != nil {
		return false
	}

	session := sessions.Default(c)
	userid := cookie.Value
	userJSON := session.Get(userid)
	if userJSON == nil {
		return false
	}

	var userInfo models.UserInfo
	json.Unmarshal([]byte(userJSON.(string)), &userInfo)

	return time.Since(userInfo.LoginTime) <= 5*time.Minute
}
