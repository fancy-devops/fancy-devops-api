package main

import (
	"fmt"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/fancy-devops/fancy-devops-api/docs"
	"github.com/fancy-devops/fancy-devops-api/routers"
	"github.com/fancy-devops/fancy-devops-api/utils/setting"
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
