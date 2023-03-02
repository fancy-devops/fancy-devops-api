package main

import (
	"fmt"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gitlab.chad122.top/fancy-devops/fancy-devops-api/docs"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/setting"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/routers"
)

func main() {
	router := routers.InitRouter()

	// swag init
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
