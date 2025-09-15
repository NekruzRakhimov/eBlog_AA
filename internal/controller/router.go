package controller

import (
	_ "eBlog/docs"
	"eBlog/internal/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (ctrl *Controller) InitRoutes() error {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", ctrl.ping)

	authG := r.Group("/auth")
	authG.POST("/sign-up", ctrl.SignUp)
	authG.POST("/sign-in", ctrl.SignIn)
	authG.GET("/refresh", ctrl.RefreshTokenPair)

	apiGroup := r.Group("/api", ctrl.checkUserAuthentication)

	apiGroup.POST("/articles", ctrl.createArticle) // /api/articles   /articles -> 404
	apiGroup.GET("/articles", ctrl.getAllArticles)
	apiGroup.GET("/articles/:id", ctrl.getArticleByID)
	apiGroup.PUT("/articles/:id", ctrl.updateArticle)
	apiGroup.DELETE("/articles/:id", ctrl.deleteArticle)

	err := r.Run(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun))
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *Controller) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
