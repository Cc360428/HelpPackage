package main

import (
	"github.com/Cc360428/HelpPackage/gin_utils"
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/auth", Auth)
	r.GET("/get/one", GetOne)
	_ = r.Run(":8888")
}

type Login struct {
	Username string `json:"username" `
	Password string `json:"password"`
}

// Auth 鉴权
func Auth(context *gin.Context) {
	var request Login
	err := gin_utils.ParseJSON(context, &request)
	logs.Info(request)
	if err != nil {
		gin_utils.ResponseErrorBody(context, err)
	} else {
		gin_utils.ResponseSuccessBody(context, request)
	}
}

// GetOne 获取单个
func GetOne(context *gin.Context) {
	param := context.Query("id")

	logs.Info("参数：", param)
}
