package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 400
	SUCCESS = 200
)

// Success success result
func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code:    SUCCESS,
		Message: "success",
		Data:    data,
	})
}

// Fail failed result
func Fail(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Result{
		Code:    ERROR,
		Message: message,
		Data:    nil,
	})
}

func Failed(code int, message string, c *gin.Context) {
	c.JSON(code, Result{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func SuccessWithPagination(data any, p models.Pagination, c *gin.Context) {
	c.JSON(http.StatusOK, models.ApiResponse{
		Code:       0,
		Message:    "success",
		Data:       data,
		Pagination: p,
	})
}
