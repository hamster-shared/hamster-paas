package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/rpc/aline"
	"log"
	"net/http"
	"os"
	"strings"
)

func (h *HandlerServer) Authorize() gin.HandlerFunc {
	return func(gin *gin.Context) {

		var userPrincipal aline.UserPrincipal
		jwtToken := gin.GetHeader("Authorization")
		log.Println(jwtToken)
		if jwtToken == "" {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			gin.Abort()
			return
		}
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		userId, ok := claims["userId"].(float64)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		loginType, ok := claims["loginType"].(float64)
		if !ok {
			Failed(http.StatusUnauthorized, "Invalid token", gin)
			gin.Abort()
			return
		}
		log.Println(loginType)
		gin.Set("loginType", int(loginType))
		githubToken := ""
		userService, err := application.GetBean[*aline.UserService]("userService")
		if loginType == consts.GitHub {
			user, err := userService.GetUserById(int64(userId))
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				gin.Abort()
				return
			}
			githubToken = user.Token
			gin.Set("user", user)
			userPrincipal = &user
			gin.Set("userPrincipal", userPrincipal)
		}
		if loginType == consts.Metamask {
			log.Println(userId)
			userWallet, err := userService.GetUserWalletById(int(userId))
			log.Println(userWallet)
			if err != nil {
				Failed(http.StatusUnauthorized, err.Error(), gin)
				gin.Abort()
				return
			}
			gin.Set("user", userWallet)
			userPrincipal = &userWallet
			gin.Set("userPrincipal", userPrincipal)
		}
		gin.Set("token", githubToken)
		gin.Next()
	}
}
