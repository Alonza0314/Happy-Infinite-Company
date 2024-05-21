package auth

import (
	"encoding/json"
	"hic/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostSignup(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	client := models.NewClient(username, email, password)

	var redirectURL string
	if err := models.ProcessSignup(client); err != nil {
		redirectURL = "/login/?signup=" + url.QueryEscape(err.Error())
	} else {
		redirectURL = "/login/?signup=true"
	}
	c.Redirect(http.StatusFound, redirectURL)
}

func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	client := models.NewClient(username, "", password)

	if err := models.ProcessLogin(client); err != nil {
		c.Redirect(http.StatusFound, "/login/?login="+url.QueryEscape(err.Error()))
		c.Abort()
		return
	}

	// set session
	userInfo := models.UserInfo{Username: username, LoginTime: time.Now()}
	userJSON, err := json.Marshal(userInfo)
	if err != nil {
		c.Redirect(http.StatusFound, "/login/?login="+url.QueryEscape("服務器json錯誤"))
		c.Abort()
		return
	}
	session := sessions.Default(c)
	userid := models.GenerateHash(userInfo)
	session.Set(userid, string(userJSON))
	session.Save()

	// set cookie "userid"
	c.SetCookie("userid", userid, 5*60, "/", "", false, true)
	/*
		1. name：cookie 的名称。
		2. value：cookie 的值。
		3. maxAge：cookie 的过期时间（以秒为单位），如果为正数，则表示 cookie 在指定的秒数后过期；如果为负数，则表示 cookie 在浏览器关闭后过期；如果为 0，则表示立即删除该 cookie。在示例中，3600 表示 cookie 在 3600 秒（即 1 小时）后过期。
		4. path：cookie 的作用路径。指定 cookie 生效的路径，浏览器只会向该路径发送 cookie。在示例中，"/" 表示整个网站都可以接收到该 cookie。
		5. domain：cookie 的作用域。指定哪些域名可以接收到 cookie。如果为空字符串，则表示只有设置该 cookie 的域名可以接收到；如果为 nil，则表示使用当前请求的域名。在示例中，"" 表示使用当前请求的域名。
		6. secure：是否只在 HTTPS 连接下传输 cookie。如果为 true，则表示只在 HTTPS 连接下传输；如果为 false，则表示在 HTTP 和 HTTPS 连接下都传输。在示例中，false 表示在 HTTP 和 HTTPS 连接下都传输。
		7. httpOnly：是否限制 cookie 只能通过 HTTP 或 HTTPS 协议传输，而不能通过 JavaScript 访问。如果为 true，则表示只能通过 HTTP 或 HTTPS 协议传输；如果为 false，则表示可以通过 JavaScript 访问。在示例中，true 表示只能通过 HTTP 或 HTTPS 协议传输。
	*/

	c.Redirect(http.StatusFound, "/subpage")
}

func PostFindpw(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	client := models.NewClient(username, email, "")

	// 確認email
	if err := models.ProcessFindpw(client); err != nil {
		c.Redirect(http.StatusFound, "/login/findpw/?findpw="+url.QueryEscape(err.Error()))
		c.Abort()
		return
	}

	// 設定session
	resetInfo := models.ResetInfo{Username: username, Email: email}
	resetJSON, err := json.Marshal(resetInfo)
	if err != nil {
		c.Redirect(http.StatusFound, "/login/findpw/?findpw="+url.QueryEscape("服務器json錯誤"))
		c.Abort()
		return
	}
	session := sessions.Default(c)
	session.Set("resetid", string(resetJSON))
	session.Save()

	c.Redirect(http.StatusFound, "/login/resetpw")
}

func PostResetpw(c *gin.Context) {
	password := c.PostForm("password")

	// 取得session
	session := sessions.Default(c)
	resetJSON := session.Get("resetid")
	var resetInfo models.ResetInfo
	json.Unmarshal([]byte(resetJSON.(string)), &resetInfo)

	client := models.NewClient(resetInfo.Username, resetInfo.Email, password)

	// 重設密碼
	if err := models.ProcessResetpw(client); err != nil {
		c.Redirect(http.StatusFound, "/login/findpw/?findpw="+url.QueryEscape(err.Error()))
		c.Abort()
		return
	}

	c.Redirect(http.StatusFound, "/login/?resetpw="+url.QueryEscape("密碼更改成功，請重新登入"))
}
