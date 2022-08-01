// IP地址
package netc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

// 获取外网IP地址
func GetExternal() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)
	return string(content)
}

// ResolveIPV4address 解析IPv4
// 国家、地区、城市、运行商、错误
func ResolveIPV4address(ip string) (country, area, city, Isp string, err error) {
	netIp := net.ParseIP(ip)
	if netIp == nil {
		return
	} else {
		result := TaoBaoAPI(ip)
		if result != nil {
			//fmt.Println("国家：", result.Country)
			//fmt.Println("地区：", result.Timezone)
			//fmt.Println("城市：", result.City)
			//fmt.Println("运营商：", result.Isp)
			return result.Country, result.Timezone, result.City, result.Isp, err
			//return result.Data.Country, result.Data.Area, result.Data.City, result.Data.Isp, err
		}
	}
	return
}

func TaoBaoAPI(ip string) *IpApiCom {
	//http://ip.taobao.com/service/getIpInfo.php?ip=[ip地址字串]
	url := "http://ip-api.com/json/" + ip + "?lang=zh-CN"
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result IpApiCom
	if err := json.Unmarshal(out, &result); err != nil {
		return nil
	}
	return &result
}

type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}
type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

type IpApiCom struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

// 获取本地IP地址
func GetLocalIPAddress() string {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range adders {
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String()
			}
		}
	}
	return ""
}
