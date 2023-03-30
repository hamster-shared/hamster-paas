package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils"
	"net/http"
)

func (h *HandlerServer) Authorize() gin.HandlerFunc {
	return func(gin *gin.Context) {
		accessToken := gin.Request.Header.Get("Access-Token")
		if accessToken == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			gin.Abort()
			return
		}
		token := utils.AesDecrypt(accessToken, consts.SecretKey)
		userService, err := application.GetBean[*aline.UserService]("userService")
		if err != nil {
			Failed(http.StatusUnauthorized, "get user failed", gin)
			return
		}
		user, err := userService.GetUserByToken(token)
		if err != nil {
			Failed(http.StatusUnauthorized, err.Error(), gin)
			gin.Abort()
			return
		}
		if user.Token == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			gin.Abort()
			return
		}
		user.Token = accessToken
		gin.Set("token", token)
		gin.Set("user", user)
		gin.Next()
	}
}
