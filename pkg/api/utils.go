package api

import (
	"hamster-paas/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, models.ApiResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithPagination(c *gin.Context, data any, p models.Pagination) {
	c.JSON(http.StatusOK, models.ApiResponse{
		Code:       0,
		Message:    "success",
		Data:       data,
		Pagination: p,
	})
}

func Fail(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, models.ApiResponse{
		Code:    -1,
		Message: message,
	})
}

func Failed(c *gin.Context, httpCode int, errCode int, message string) {
	c.JSON(httpCode, models.ApiResponse{
		Code:    errCode,
		Message: message,
	})
}
