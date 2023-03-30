package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/rpc/aline"
	"log"
)

func (h *HandlerServer) getSubscriptionOverview(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)

	network := c.Query("network")
	if network == "" {
		Fail("network not valid", c)
		return
	}

	ov, err := h.chainLinkSubscriptionService.GetSubscriptionOverview(user.Id, network)
	if err != nil {
		log.Println(err)
		Fail(err.Error(), c)
		return
	}

	Success(ov, c)
}

func (h *HandlerServer) getSINA(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)

	sinas := h.chainLinkSubscriptionService.GetSINAByUserId(user.Id)
	Success(sinas, c)
}
