package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "user_service/api/docs"
	"user_service/api/handler"
)

// @title Authorazation
// @version 1.0
// @description API Gateway of Authorazation
// @host localhost:8085
// BasePath: /
func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := router.Group("/user")
	user.POST("/register", hand.Register)
	user.POST("/login", hand.Login)
	user.POST("/refresh", hand.RefreshToken)
	return router
}
