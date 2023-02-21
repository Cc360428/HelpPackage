/**
 * @Author: Cc
 * @Description: 获取build时间
 * @File: version
 * @Version: 1.0.0
 * @Date: 2023/2/13 10:42
 * @Software : GoLand
 */

package other

/*
const char* build_time(void)
{
    static const char* psz_build_time = "["__DATE__ "  " __TIME__ "]";
    return psz_build_time;
}
*/
import "C"
import "fmt"

const version = "v0.0.1"

func GetVersion() string {
	return fmt.Sprintf("%v %v", version, C.GoString(C.build_time()))
}
