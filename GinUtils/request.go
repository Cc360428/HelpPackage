// Gin request 请求参数
package GinUtils

// ParseJSON 获取json 和解析参数
func ParseJSON(c *gin.Context, obj interface{}) error {
	//  解析是否是json 格式
	if c.ShouldBind(&obj) == nil {
	} else {
		logs.Error("传入参数失败")
		return errors.New("解析请求参数发生错误")
	}
	return nil
}
