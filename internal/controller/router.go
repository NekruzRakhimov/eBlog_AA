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

func InitRoutes() error {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", ping)

	authG := r.Group("/auth")
	authG.POST("/sign-up", SignUp)
	authG.POST("/sign-in", SignIn)
	authG.GET("/refresh", RefreshTokenPair)

	apiGroup := r.Group("/api", checkUserAuthentication)

	apiGroup.POST("/articles", createArticle) // /api/articles   /articles -> 404
	apiGroup.GET("/articles", getAllArticles)
	apiGroup.GET("/articles/:id", getArticleByID)
	apiGroup.PUT("/articles/:id", updateArticle)
	apiGroup.DELETE("/articles/:id", deleteArticle)

	err := r.Run(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun))
	if err != nil {
		return err
	}

	return nil
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
