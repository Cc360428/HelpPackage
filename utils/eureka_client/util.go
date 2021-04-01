package eureka_client

import (
	"io/ioutil"
	"net/http"
)

// 获取公网IP
func GetLocalIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)

}

/**
 *
 源码是这样的 因 获取本地IP 不是公网的IP 服务器会IP地址错误
 // getLocalIP 获取本地ip
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
*/
