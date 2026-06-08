/**
 * @Author: Cc
 * @Description: 获取build时间
 * @File: version
 * @Version: 1.0.0
 * @Date: 2023/2/13 10:42
 * @Software : GoLand
 */

package other

import "fmt"

const version = "v0.0.1"

var buildTime = "unknown"
var buildCommit = "unknown"

func GetVersion() string {
	return fmt.Sprintf("%v %v", version, buildTime)
}

func SetBuildInfo(time, commit string) {
	buildTime = time
	buildCommit = commit
}
