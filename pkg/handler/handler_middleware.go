package handler

import "github.com/gin-gonic/gin"

func (h *HandlerServer) middlewareRpc(c *gin.Context) {
	result, err := h.middleWareService.MiddleWareRpc()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(result, c)
}
