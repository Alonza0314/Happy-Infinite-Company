package routes

import (
	"hic/api"
	"hic/auth"
	"hic/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	{
		router.GET("/", handlers.GetHome)
		aboutGroup := router.Group("/about")
		{
			aboutGroup.GET("/", handlers.GetAbout)
			aboutGroup.GET("/history", handlers.GetHistory)
			aboutGroup.GET("/members", handlers.GetMembers)
		}
		router.GET("/nothing", handlers.GetNothing)
		router.GET("/contact", handlers.GetContact)
		loginGroup := router.Group("/login")
		{
			loginGroup.GET("/", handlers.GetLogin)
			loginGroup.GET("/findpw", handlers.GetFindpw)
			loginGroup.GET("/resetpw", handlers.GetResetpw)
		}
		subpageGroup := router.Group("/subpage")
		{
			subpageGroup.GET("/", handlers.GetSubpage)
		}

		authGroup := router.Group("/api")
		{
			authGroup.GET("/captcha", api.GetCaptcha)
		}
	}

	{
		loginGroup := router.Group("/login")
		{
			loginGroup.POST("/", auth.PostLogin)
			loginGroup.POST("/signup", auth.PostSignup)
			loginGroup.POST("/findpw", auth.PostFindpw)
			loginGroup.POST("/resetpw", auth.PostResetpw)
		}
	}
}
