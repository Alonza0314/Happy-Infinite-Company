package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 返回頁面

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func GetAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}

func GetHistory(c *gin.Context) {
	c.HTML(http.StatusOK, "history.html", nil)
}

func GetMembers(c *gin.Context) {
	c.HTML(http.StatusOK, "members.html", nil)
}

func GetNothing(c *gin.Context) {
	c.HTML(http.StatusOK, "nothing.html", nil)
}

func GetContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", nil)
}

func GetLogin(c *gin.Context) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("configs/config.conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error in reading configuration file:", err)
		c.HTML(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"ActionURLSignup": "http://" + viper.GetString("httpserver.addr") + "/login/signup",
		"ActionURLLogin":  "http://" + viper.GetString("httpserver.addr") + "/login",
	})
}

func GetFindpw(c *gin.Context) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("configs/config.conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error in reading configuration file:", err)
		c.HTML(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	c.HTML(http.StatusOK, "findpw.html", gin.H{
		"ActionURLFindpw": "http://" + viper.GetString("httpserver.addr") + "/login/findpw",
		"ActionURLCaptcha": "http://" + viper.GetString("httpserver.addr") + "/api/captcha",
	})
}

func GetResetpw(c *gin.Context) {
	viper.SetConfigType("toml")
	viper.SetConfigFile("configs/config.conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error in reading configuration file:", err)
		c.HTML(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	c.HTML(http.StatusOK, "resetpw.html", gin.H{
		"ActionURLResetpw": "http://" + viper.GetString("httpserver.addr") + "/login/resetpw",
	})
}

func GetSubpage(c *gin.Context) {
	c.HTML(http.StatusOK, "subpage.html", nil)
}
