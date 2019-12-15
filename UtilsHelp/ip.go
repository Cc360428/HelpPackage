package UtilsHelp

import(
	"net"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
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



// 解析IPv4
func ResolveIPV4address (ip string)(country,area,city,Isp string,err error){
	netIp := net.ParseIP(ip)
	if netIp == nil{
		return 
	}else{
		result := TaoBaoAPI(ip)
		if result != nil {
			// fmt.Println("国家：", result.Data.Country)
			// fmt.Println("地区：", result.Data.Area)
			// fmt.Println("城市：", result.Data.City)
			// fmt.Println("运营商：", result.Data.Isp)
			return result.Data.Country,result.Data.Area,result.Data.City,result.Data.Isp ,err
		}
	}
	return 
}


func TaoBaoAPI(ip string) *IPInfo {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	url += ip

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result IPInfo
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