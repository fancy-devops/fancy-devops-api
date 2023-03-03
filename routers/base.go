package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/fancy-devops/fancy-devops-api/controller"
	"github.com/fancy-devops/fancy-devops-api/middleware/jwt"
	"github.com/fancy-devops/fancy-devops-api/utils/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	userNoAuthApis := r.Group("/api/user")
	{
		userNoAuthApis.POST("/sendverifyemail", controller.NewUser().SendVerifyEmail)
		userNoAuthApis.POST("/resetpwd", controller.NewUser().ResetPwd)
		userNoAuthApis.POST("/register", controller.NewUser().Register)
		userNoAuthApis.POST("/login", controller.NewUser().Login)
		userNoAuthApis.POST("/glogin", controller.NewUser().LoginByGoogle)
	}

	userAuthApis := r.Group("/api/user").Use(jwt.JWT())
	{
		userAuthApis.POST("/changepwd", controller.NewUser().ChangePwd)
		userAuthApis.GET("/list", controller.NewUser().GetUserList)
		userAuthApis.GET("/detail/:id", controller.NewUser().GetUserDetail)
	}

	gitNoAuthApis := r.Group("/api/git")
	{
		gitNoAuthApis.POST("/hook", controller.NewGit().Hook)
	}

	return r
}
