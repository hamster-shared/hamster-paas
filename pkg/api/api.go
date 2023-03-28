package api

import (
	"fmt"
	"hamster-paas/pkg/models"
	"os"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.GET("/chains", chains)
	r.GET("/networks/:chain", networks)
	r.GET("/apps/:account", getApps)
	r.POST("/app", createApp)
	r.DELETE("/app/:account/:app_id", deleteApp)

	r.GET("/subscription/overview", subscriptionOverview)

	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func chains(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "chains",
	})
}

func networks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "networks",
	})
}

func getApps(c *gin.Context) {
	// account, ok := c.Params.Get("account")
	// if !ok {
	// 	Fail(c, "invalid params")
	// 	return
	// }
	// var pagination models.ApiRequestPagination
	// if err := c.ShouldBindQuery(&pagination); err != nil {
	// 	Fail(c, "invalid params")
	// 	return
	// }

	c.JSON(200, gin.H{
		"message": "apps",
	})
}

func createApp(c *gin.Context) {
	var appParams models.ApiRequestCreateApp
	if err := c.ShouldBindJSON(&appParams); err != nil {
		Fail(c, "invalid params")
		return
	}
	// TODO: create app

	c.JSON(200, gin.H{
		"message": "createApp",
	})
}

func deleteApp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleteApp",
	})
}

func subscriptionOverview(c *gin.Context) {

}


