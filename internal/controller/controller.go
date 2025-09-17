package controller

import (
	"eBlog/internal/contracts"
	"eBlog/internal/errs"
	"eBlog/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service contracts.ServiceI
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		service: svc,
	}
}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrNotFound) || errors.Is(err, errs.ErrArticleNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidRequestBody) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists) ||
		errors.Is(err, errs.ErrInvalidPathParam):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrFillRequiredFields):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword) || errors.Is(err, errs.ErrUserIDNotFoundInContext):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
