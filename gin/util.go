package gin_utils

//import

// ParseJSON 解析请求JSON

// Wrap400Response 包装错误码为400的响应错误
//func Wrap400Response(err error, msg ...string) error {
//	m := "请求发生错误"
//	if len(msg) > 0 {
//		m = msg[0]
//	}
//	return WrapResponse(err, 400, m, 400)
//}

// WrapResponse 包装响应错误
//func WrapResponse(err error, code int, msg string, status ...int) error {
//	res := &ResponseError{
//		Code:    code,
//		Message: msg,
//		ERR:     err,
//	}
//	if len(status) > 0 {
//		res.StatusCode = status[0]
//	}
//	return res
//}
