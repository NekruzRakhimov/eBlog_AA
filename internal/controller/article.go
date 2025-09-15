package controller

import (
	"eBlog/internal/models"
	"eBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateArticleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// createArticle
// @Summary     Создать новую статью
// @Description Роут для создания новой статьи
// @Tags        Articles
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body CreateArticleRequest true "данные для создания новой статьи"
// @Success     201 {object} CommonResponse
// @Failure     401 {object} CommonError
// @Failure     400 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles [post]
func createArticle(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, CommonError{"user id not found in context"})
		return
	}

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, CommonError{err.Error()})
		return
	}

	input.UserID = userID

	if err := service.CreateArticle(input); err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{"article created successfully"})
}

// getAllArticles
// @Summary     Получение списка всех статей
// @Description Получения списка статей которые принадлежат текущему пользователю
// @Tags        Articles
// @Produce     json
// @Param       title query string false "поиск по title"
// @Security    BearerAuth
// @Success     201 {array} models.Article
// @Failure     401 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles [get]
func getAllArticles(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in context"})
		return
	}

	title := c.Query("title")

	articles, err := service.GetAllArticles(userID, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, articles)
}

// getArticleByID
// @Summary     Получение статьи по ID
// @Description Получение статьи по ID которая принадлежит текущему пользователю
// @Tags        Articles
// @Produce     json
// @Param       id path int true "id статьи"
// @Security    BearerAuth
// @Success     201 {object} models.Article
// @Failure     401 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles/{id} [get]
func getArticleByID(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in context"})
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID статьи должен быть числом",
		})
		return
	}

	article, err := service.GetArticleByID(id)
	if err != nil {
		if err.Error() == "статья c таким ID не найдена" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, article)
}

// updateArticle
// @Summary     Обновление статьи по ID
// @Description Обновление статьи по ID которая принадлежит текущему пользователю
// @Tags        Articles
// @Produce     json
// @Param       id path int true "id статьи"
// @Param       request body CreateArticleRequest true "данные для обновления статьи"
// @Security    BearerAuth
// @Success     201 {object} models.Article
// @Failure     401 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles/{id} [put]
func updateArticle(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in context"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) //f
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID статьи должен быть числом"})
		return
	}

	var input models.Article

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "заголовок статьи не может быть пустым"})
		return
	}

	if input.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "описание статьи не может быть пустым"})
		return
	}

	input.ID = id

	if err = service.UpdateArticle(input); err != nil {
		if err.Error() == "статья c таким ID не найдена" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "статья успешно обновлена"})
}

// deleteArticle
// @Summary     Удаление статьи по ID
// @Description Удаление статьи по ID которая принадлежит текущему пользователю
// @Tags        Articles
// @Produce     json
// @Param       id path int true "id статьи"
// @Security    BearerAuth
// @Success     201 {object} models.Article
// @Failure     401 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles/{id} [delete]
func deleteArticle(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in context"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID статьи должен быть числом"})
		return
	}

	if err = service.DeleteArticle(id); err != nil {
		if err.Error() == "статья c таким ID не найдена" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "статья успешно удалена"})
}
