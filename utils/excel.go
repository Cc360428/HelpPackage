package utils

import "github.com/Cc360428/HelpPackage/utils/logs"

//Excel导出
/**
1、传入 []*类型
2、返回下载路径
3、提供之后删除文件
*/

//Excel导出
func Excel(obj interface{}) (downloadPath string, err error) {
	logs.Info(obj)
	return "s", err
}
