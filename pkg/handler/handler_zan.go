package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service/zan"
	"strconv"
)

var TimeIntervalLimit = []string{"STAT_15_MIN", "STAT_1_HOUR", "STAT_24_HOUR", "STAT_7_DAY", "STAT_1_MONTH"}

// ZanAuthed godoc
// @Security ApiKeyAuth
// @Summary 用户是否已经进行了zan授权
// @Schemes
// @Description 判断用户是否已经进行了zan授权，如果未授权，需要进行授权流程
// @description 授权流程
// @description        1. 调用 获取授权链接接⼝ /api/v2/zan/auth_url
// @description        2. 前端获取到zan平台跳转的url ，进行跳转
// @description        3. 用户在zan平台进行登陆和授权行为
// @description        4. zan平台将浏览器url跳回到hamster平台，地址为：
// @description        5. 前端通过url解析出authCode参数，调用 交换zan的access_token 接口 /api/v2/zan/exchange/access_token
// @description        6. 调用其他zan平台相关需要授权的接口，如创建ApiKEY
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=bool} "desc"
// @Router /api/v2/zan/account/authed [get]
func (h *HandlerServer) ZanAuthed(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}

	authed := h.zanService.GetUserAuthed(u)
	Success(authed, c)

}

// ZanGetAuthUrl godoc
// @Security ApiKeyAuth
// @Summary 获取授权链接接⼝
// @Schemes
// @Description 获取授权链接接⼝
// @description 返回结果中的data 是前端需要跳转到zan平台进行认证的url
// @description zan平台认证完成后，会跳转回前端，如
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result
// @Router /api/v2/zan/account/auth_url [get]
func (h *HandlerServer) ZanGetAuthUrl(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	result, err := h.zanService.GetAuthUrl()

	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(result, c)
}

// ZanExchangeAccessToken godoc
// @Security ApiKeyAuth
// @Summary 交换zan的access_token
// @Schemes
// @Description 用户在zan平台授权hamster访问后，zan平台通过跳转方式返回给前端authCode, 此时需要调用后端接口与zan平台交换成可以访问的access_token
// @description 前端需要从url中解析这个authCode，并调用此接口交换
// @description 用户第一次交换成功后，后续不需要再此交换，此token由后台保存
// @Param exchangeAccessTokenVo body vo.ExchangeAccessTokenVo true "请求交换token参数"
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result
// @Router /api/v2/zan/account/access_token [post]
func (h *HandlerServer) ZanExchangeAccessToken(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	var req vo.ExchangeAccessTokenVo

	err := c.BindJSON(&req)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	u, ok := user.(aline.User)

	if !ok {
		Fail("cannot get user info", c)
		return
	}
	err = h.zanService.ExchangeAccessToken(u, req.AuthCode)

	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success("", c)
}

// ZanCreateApiKey godoc
// @Security ApiKeyAuth
// @Summary 创建API KEY接⼝
// @Schemes
// @Description 为⽤户创建API KEY
// @Param apiKeyCreateReq body zan.ApiKeyCreateReq true "创建api key"
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=zan.ApiKeyBase}
// @Router /api/v2/zan/node-service/api-keys [post]
func (h *HandlerServer) ZanCreateApiKey(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	var req zan.ApiKeyCreateReq

	err := c.BindJSON(&req)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	u, ok := user.(aline.User)

	if !ok {
		Fail("cannot get user info", c)
		return
	}
	resp, err := h.zanService.CreateApiKey(u, req)

	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanApiKeyPage godoc
// @Security ApiKeyAuth
// @Summary API KEY分⻚查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝分⻚查询⽤户的API KEY列表
// @Param page query int true "页码" default(1)
// @Param size query int true "每⻚数量" default(10)
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=vo.PageResp[zan.ApiKeyDigestInfo]}
// @Router /api/v2/zan/node-service/api-keys/list [get]
func (h *HandlerServer) ZanApiKeyPage(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	p := c.DefaultQuery("page", "1")
	s := c.DefaultQuery("size", "10")

	page, err := strconv.Atoi(p)
	if err != nil {
		Fail("param page is not int", c)
		return
	}
	size, err := strconv.Atoi(s)
	if err != nil {
		Fail("param size is not int", c)
		return
	}
	u, ok := user.(aline.User)

	if !ok {
		Fail("cannot get user info", c)
		return
	}
	resp, err := h.zanService.ApiKeyList(u, page, size)

	if err != nil {
		fmt.Println(err)
		Fail(err.Error(), c)
		return
	}

	pageResponse := vo.PageResp[zan.ApiKeyDigestInfo]{
		Total:     resp.Total,
		PageCount: resp.PageCount,
		Page:      resp.PageNum,
		PageSize:  resp.PageSize,
		Data:      resp.Data,
	}

	Success(pageResponse, c)
}

// ZanApiKeyDetail godoc
// @Security ApiKeyAuth
// @Summary API KEY详情接口
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询特定API KEY的详情
// @Param apiKeyId query string true "apiKeyId"
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=zan.ApiKeyDetailInfo}
// @Router /api/v2/zan/node-service/api-keys/detail [get]
func (h *HandlerServer) ZanApiKeyDetail(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	apiKeyId := c.DefaultQuery("apiKeyId", "")
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}
	resp, err := h.zanService.ApiKeyDetail(u, apiKeyId)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanApiKeyCreditCost godoc
// @Security ApiKeyAuth
// @Summary API KEY credit cost统计查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询当前API KEY在过去24⼩时credit消耗量统计数据
// @Param apiKeyId query string true "apiKeyId"
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=[]zan.StatCreditCostItem}
// @Router /api/v2/zan/node-service/api-keys/stats/credit-cost [get]
func (h *HandlerServer) ZanApiKeyCreditCost(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	apiKeyId := c.DefaultQuery("apiKeyId", "")
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}
	resp, err := h.zanService.ApiKeyCreditCost(u, apiKeyId)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanApiKeyRequestStats godoc
// @Security ApiKeyAuth
// @Summary API Key request统计查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询当前API KEY过去⼀段时间不同⽣态下的不同⽅法的接⼝调⽤统计数据
// @Param apiKeyId query string true "apiKeyId"
// @Param timeInterval query string true "时间间隔"  Enums(STAT_15_MIN,STAT_1_HOUR,STAT_24_HOUR,STAT_7_DAY,STAT_1_MONTH)
// @Param ecosystem query string true "⽣态, 此值为链⽣态摘要信息查询接⼝返回的生态编码"
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=[]zan.StatMethodCountItem}
// @Router /api/v2/zan/node-service/api-keys/stats/requests [get]
func (h *HandlerServer) ZanApiKeyRequestStats(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	apiKeyId := c.DefaultQuery("apiKeyId", "")
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}
	timeInterval := c.DefaultQuery("timeInterval", "STAT_15_MIN")
	if !lo.Contains[string](TimeIntervalLimit, timeInterval) {
		Fail("timeInterval is not in [STAT_15_MIN,STAT_1_HOUR,STAT_24_HOUR,STAT_7_DAY,STAT_1_MONTH]", c)
		return
	}
	ecosystem := c.GetString("ecosystem")
	if ecosystem == "" {
		Fail("param ecosystem is required", c)
		return
	}

	resp, err := h.zanService.ApiKeyRequestStats(u, apiKeyId, timeInterval, ecosystem)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanApiKeyRequestActivityStats godoc
// @Security ApiKeyAuth
// @Summary API KEY requests activity统计查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询当前API KEY过去⼀段时间不同⽣态下不同⽅法的接⼝调⽤成功次数统计数据
// @Param apiKeyId query string true "apiKeyId"
// @Param timeInterval query string true "时间间隔"  Enums(STAT_15_MIN,STAT_1_HOUR,STAT_24_HOUR,STAT_7_DAY,STAT_1_MONTH)
// @Param ecosystem query string true "⽣态, 此值为链⽣态摘要信息查询接⼝返回的生态编码"
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=[]zan.StatMethodRequestActivityDetail}
// @Router /api/v2/zan/node-service/api-keys/stats/requests-activity [get]
func (h *HandlerServer) ZanApiKeyRequestActivityStats(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	apiKeyId := c.DefaultQuery("apiKeyId", "")
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}
	timeInterval := c.DefaultQuery("timeInterval", "STAT_15_MIN")
	if !lo.Contains[string](TimeIntervalLimit, timeInterval) {
		Fail("timeInterval is not in [STAT_15_MIN,STAT_1_HOUR,STAT_24_HOUR,STAT_7_DAY,STAT_1_MONTH]", c)
		return
	}
	ecosystem := c.GetString("ecosystem")
	if ecosystem == "" {
		Fail("param ecosystem is required", c)
		return
	}

	resp, err := h.zanService.ApiKeyRequestActivityStats(u, apiKeyId, timeInterval, ecosystem)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanApiKeyRequestOriginStats godoc
// @Security ApiKeyAuth
// @Summary API KEY request Origin统计查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询当前API KEY过去⼀段时间不同请求来源的统计数据
// @Param apiKeyId query string true "apiKeyId"
// @Param timeInterval query string true "时间间隔"  Enums(STAT_15_MIN,STAT_1_HOUR,STAT_24_HOUR,STAT_7_DAY,STAT_1_MONTH)
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=[]zan.StatCreditCostOrigin}
// @Router /api/v2/zan/node-service/api-keys/stats/requests-origin [get]
func (h *HandlerServer) ZanApiKeyRequestOriginStats(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}

	apiKeyId := c.DefaultQuery("apiKeyId", "")
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}
	timeInterval := c.DefaultQuery("timeInterval", "STAT_15_MIN")
	if !lo.Contains[string](TimeIntervalLimit, timeInterval) {
		Fail("timeInterval is not in [STAT_15_MIN,STAT_1_HOUR,STAT_24_HOUR,STAT_7_DAY,STAT_1_MONTH]", c)
		return
	}

	resp, err := h.zanService.ApiKeyRequestOriginStats(u, apiKeyId, timeInterval)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanEcosystemsDigest godoc
// @Security ApiKeyAuth
// @Summary 链⽣态摘要信息查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询特定API KEY的详情。
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=[]zan.EcosystemDigestInfo}
// @Router /api/v2/zan/node-service/ecosystems/digest [get]
func (h *HandlerServer) ZanEcosystemsDigest(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}

	resp, err := h.zanService.EcosystemsDigest(u)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}

// ZanPlan godoc
// @Security ApiKeyAuth
// @Summary 套餐信息查询接⼝
// @Schemes
// @Description 调⽤⽅可通过该接⼝查询当前⽤户的节点服务套餐情况
// @Tags zan
// @Accept json
// @Produce json
// @Success 200 {object} Result{data=[]zan.EcosystemDigestInfo}
// @Router /api/v2/zan/node-service/plan [get]
func (h *HandlerServer) ZanPlan(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	u, ok := user.(aline.User)
	if !ok {
		Fail("cannot get user info", c)
		return
	}

	resp, err := h.zanService.UserPlan(u)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(resp, c)
}
