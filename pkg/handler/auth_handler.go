package handler

import (
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *HandlerServer) Authorize() gin.HandlerFunc {
	return func(gin *gin.Context) {

		var userPrincipal aline.UserPrincipal
		jwtToken := gin.GetHeader("Authorization")
		log.Println(jwtToken)
		if jwtToken == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			logger.Errorf("Authorization: access not authorized")
			gin.Abort()
			return
		}
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)
		token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		// if err != nil || !token.Valid {
		// 	Failed(http.StatusUnauthorized, "Invalid token", gin)
		// 	logger.Errorf("Authorization: Invalid token 1")
		// 	gin.Abort()
		// 	return
		// }
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			logger.Errorf("Authorization: Invalid token 2")
			gin.Abort()
			return
		}
		userId, ok := claims["userId"].(float64)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			logger.Errorf("Authorization: Invalid token 3")
			gin.Abort()
			return
		}
		loginType, ok := claims["loginType"].(float64)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			logger.Errorf("Authorization: Invalid token 4")
			gin.Abort()
			return
		}
		logger.Infof("Authorization: userId: %d, loginType: %d", int(userId), int(loginType))
		gin.Set("loginType", int(loginType))
		githubToken := ""
		userService, _ := application.GetBean[*aline.UserService]("userService")
		if loginType == consts.GitHub {
			user, err := userService.GetUserById(int64(userId))
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				logger.Errorf("Authorization: GitHub StatusUnauthorized %v", err)
				gin.Abort()
				return
			}
			logger.Infof("Authorization: user: %+v", user)
			githubToken = user.Token
			gin.Set("user", user)
			gin.Set("userId", user.Id)
			userPrincipal = &user
			gin.Set("userPrincipal", userPrincipal)
		}
		if loginType == consts.Metamask {
			log.Println(userId)
			userWallet, err := userService.GetUserWalletById(int(userId))
			log.Println(userWallet)
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				logger.Errorf("Authorization: Metamask StatusUnauthorized %v", err)
				gin.Abort()
				return
			}
			logger.Infof("Authorization: userWallet: %+v", userWallet)
			gin.Set("user", userWallet)
			gin.Set("userId", userWallet.UserId)
			userPrincipal = &userWallet
			gin.Set("userPrincipal", userPrincipal)
		}
		gin.Set("token", githubToken)
		gin.Next()
	}
}
