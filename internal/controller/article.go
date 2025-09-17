package controller

import (
	"eBlog/internal/errs"
	"eBlog/internal/models"
	"errors"
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
// @Failure     422 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles [post]
func (ctrl *Controller) createArticle(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		ctrl.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	input.UserID = userID

	if input.Title == "" || input.Description == "" {
		ctrl.handleError(c, errs.ErrFillRequiredFields)
		return
	}

	if err := ctrl.service.CreateArticle(input); err != nil {
		ctrl.handleError(c, err)
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
func (ctrl *Controller) getAllArticles(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		ctrl.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	title := c.Query("title")

	articles, err := ctrl.service.GetAllArticles(userID, title)
	if err != nil {
		ctrl.handleError(c, err)
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
func (ctrl *Controller) getArticleByID(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		ctrl.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidPathParam, err))
		return
	}

	if id < 1 {
		ctrl.handleError(c, errs.ErrInvalidPathParam)
		return
	}

	article, err := ctrl.service.GetArticleByID(id)
	if err != nil {
		ctrl.handleError(c, err)
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
// @Failure     400 {object} CommonError
// @Failure     422 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles/{id} [put]
func (ctrl *Controller) updateArticle(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		ctrl.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidPathParam, err))
		return
	}

	if id < 1 {
		ctrl.handleError(c, errs.ErrInvalidPathParam)
		return
	}

	var input models.Article

	if err = c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if input.Title == "" || input.Description == "" {
		ctrl.handleError(c, errs.ErrFillRequiredFields)
		return
	}

	input.ID = id

	if err = ctrl.service.UpdateArticle(input); err != nil {
		ctrl.handleError(c, err)
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
// @Failure     400 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /api/articles/{id} [delete]
func (ctrl *Controller) deleteArticle(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		ctrl.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidPathParam, err))
		return
	}

	if id < 1 {
		ctrl.handleError(c, errs.ErrInvalidPathParam)
		return
	}

	if err = ctrl.service.DeleteArticle(id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "статья успешно удалена"})
}
