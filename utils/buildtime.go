package utils

/*
const char* build_time(void)
{
static const char* cgo_build_time = "["__DATE__ " " __TIME__ "]";
    return cgo_build_time;
}
*/
import "C"

// BuildTime 获取编译时的时间，可以作为服务器版本号
// 格式为 [Mmm dd yyyy hh:mm:ss]
func BuildTime() string {
	return C.GoString(C.build_time())
}
