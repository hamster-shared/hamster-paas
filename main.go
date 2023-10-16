package main

import (
	"hamster-paas/pkg/initialization"
)

// @title hamster paas API 接口文档
// @version 0.0.1
// @description 提供zan相关接口
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Access-Token
// @BasePath /
func main() {
	initialization.Init()
}
