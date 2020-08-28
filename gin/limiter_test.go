package gin_utils

import (
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func TestNewRateLimiter(t *testing.T) {
	router := gin.Default() //获得路由实例
	// 设置一天只能修改三次 重启项目之后会失效计时 ，建议真正的做防控建议持久计时
	reachLimiter := NewRateLimiter(time.Hour*24, 3, func(ctx *gin.Context) (string, error) {
		return ctx.Request.Header.Get("Authorization"), nil
	})
	router.POST("/update_password", reachLimiter.Middleware(), UpdatePassword)
}

func UpdatePassword(context *gin.Context) {

}
