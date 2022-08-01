// Package beego_utils beego 获取参数
package beego_utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// HelperGetInt64Param Get 获取int64
func HelperGetInt64Param(this *beego.Controller, para string) int64 {
	pathint, err := this.GetInt64(para, 0)
	if err != nil {
		log.Printf("get int parameter err :: %v", err.Error())
		ReturnFail(this, "get int parameter err :: "+err.Error())
	}
	return pathint
}

// HelperGetIntParam Get 获取int
func HelperGetIntParam(this *beego.Controller, para string) int {
	pathint, err := this.GetInt(para, 0)
	if err != nil {
		log.Printf("get int parameter err :: %v", err.Error())
		ReturnFail(this, "get int parameter err :: "+err.Error())
	}
	return pathint
}

// HelperGetStringParam Get 获取string
func HelperGetStringParam(this *beego.Controller, para string) string {
	pathStr := this.GetString(para)
	if pathStr == "" {
		ReturnFail(this, "parameter: "+para+" is null")
	}
	return pathStr
}

// HelperGetPath Get 获取string url/:
func HelperGetPath(this *beego.Controller, path string) string {
	paths := this.GetString(":" + path)
	if paths == "" {
		ReturnFail(this, "")
	}
	return paths
}

// HelperConvectBody Post 获取json
func HelperConvectBody(this *beego.Controller, out interface{}) {
	var err error
	result := ResultInit()
	if this == nil || out == nil {
		err = fmt.Errorf("PostBody为空")
		result.Msg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		this.StopRun()
	}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, out)
	if err != nil {
		log.Printf("Body (%s) 转换错误， err = %v", string(this.Ctx.Input.RequestBody), err.Error())
		result.Msg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		this.StopRun()
	}
}

// HelperConfigSessStringcfg 获取 session key string
func HelperConfigSessStringcfg(sessionName string, configName string) (cfg string, err error) {
	// 从配置文件中获取 session 和key 对应的 配置string
	cfgSess, err := beego.AppConfig.GetSection(strings.ToLower(sessionName))
	if err != nil {
		log.Printf("session (%v) err :: %v", sessionName, err.Error())
		return "", err
	} else {
		cfg = cfgSess[strings.ToLower(configName)]
		if cfg == "" {
			err = fmt.Errorf("config (%v) of session (%v) is null ", configName, sessionName) // errors.New(fmt.Sprintf("config (%v) of session (%v) is null ", configName, sessionName))
			log.Println(err.Error())
			return "", err
		}
	}
	return cfg, nil
}

// HelperConvertMac 转换macaddr，去掉 “: - ”以及空格
func HelperConvertMac(mac *string) (err error) {

	if mac == nil {
		err = fmt.Errorf("mac is null")
		log.Println(err.Error())
		return err
	}
	*mac = strings.TrimSpace(*mac)
	*mac = strings.Replace(*mac, ":", "", 5)
	*mac = strings.Replace(*mac, "-", "", 5)
	*mac = strings.ToUpper(*mac)

	return nil
}

func HelperConvetInterface(in interface{}, out interface{}) (err error) {
	inBytes, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(inBytes, out)
}

func HelperTimeGetDaySection() (startTime time.Time, endTime time.Time, err error) {

	today := time.Now()

	stStr := today.Format("2006-01-02")
	startStr := stStr + " 00:00:00"
	startTime, err = time.Parse("2006-01-02 15:04:05", startStr)
	if err != nil {
		log.Printf("get day section err :: %v", err.Error())
		return startTime, endTime, err
	}
	endStr := stStr + " 23:59:59"
	endTime, err = time.Parse("2006-01-02 15:04:05", endStr)
	if err != nil {
		return startTime, endTime, err
	}
	startTime = startTime.Local()
	endTime = endTime.Local()

	return startTime, endTime, nil
}

type Network struct {
	Name       string
	IP         string
	Mask       string
	MACAddress string
}

//
type intfInfo struct {
	Name       string
	MacAddress string
	Ipv4       []string
}

func HelperGetNetworkInfo() ([]*Network, error) {
	var nws []*Network
	intf, err := net.Interfaces()
	if err != nil {
		log.Printf("get network info failed: %v", err.Error())
		return nil, err
	}
	var is = make([]intfInfo, len(intf))
	for i, v := range intf {
		ips, err := v.Addrs()
		if err != nil {
			log.Printf("get network addr failed: %v", err.Error())
			return nil, err
		}
		//此处过滤loopback（本地回环）和isatap（isatap隧道）
		if !strings.Contains(v.Name, "Loopback") && !strings.Contains(v.Name, "isatap") {
			var network Network
			is[i].Name = v.Name
			is[i].MacAddress = v.HardwareAddr.String()
			for _, ip := range ips {
				if strings.Contains(ip.String(), ".") {
					is[i].Ipv4 = append(is[i].Ipv4, ip.String())
				}
			}
			network.Name = is[i].Name
			network.MACAddress = is[i].MacAddress
			if len(is[i].Ipv4) > 0 {
				network.IP = is[i].Ipv4[0]
			}

			//logs.Debug("network:= %v ", network)
			if len(network.IP) > 0 && len(network.MACAddress) > 0 {
				ipMasks := strings.Split(network.IP, "/")
				network.IP = ipMasks[0]
				network.Mask = ipMasks[1]
				nws = append(nws, &network)
			}
		}
	}
	return nws, nil
}

func HelperDealMac(mac string) (ipaddr string) {
	ipaddr = mac
	ipaddr = strings.Replace(ipaddr, ":", "", 5)
	ipaddr = strings.Replace(ipaddr, "-", "", 5)
	ipaddr = strings.ToUpper(ipaddr)
	ipaddr = strings.TrimSpace(ipaddr)
	return ipaddr
}

func HelperGetNetwork() (macaddr string, ip string, err error) {
	ComputerIpPrex := beego.AppConfig.String("ComputerIpPrex")
	if ComputerIpPrex == "" {
		err = fmt.Errorf("computerIpPrex 必须配置 如 ： 192.168.3") // errors.New("ComputerIpPrex 必须配置 如 ： 192.168.3.")
		log.Println(err.Error())
		return macaddr, ip, err
	}
	networks, err := HelperGetNetworkInfo()
	if err != nil {
		log.Printf("networks warn :: %v", err.Error())
		return macaddr, ip, err
	}

	for loop := 0; loop < len(networks); loop++ {
		nk := networks[loop]
		nkIp := nk.IP
		if strings.HasPrefix(nkIp, ComputerIpPrex) {
			macaddr = HelperDealMac(nk.MACAddress)
			ip = nkIp
			return macaddr, ip, nil
		}
	}

	return macaddr, ip, errors.New("no find mac and ip")
}

// HelperFileSuffix 获取文件 扩展名
func HelperFileSuffix(filename string) (suffix string) {
	fileArray := strings.Split(filename, ".")
	if len(fileArray) < 1 {
		return suffix
	}
	return fileArray[len(fileArray)-1]
}

// HelperMd5andSalt 获取 md5 加密字符串， 有盐值 (srcStr string 未使用已删除)
func HelperMd5andSalt(salt string) (md5Str string) {
	m5 := md5.New()
	m5.Write([]byte("Mi Ma"))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	md5Str = hex.EncodeToString(st)
	return md5Str
}

// 时间转字符串
const baseFormat = "2006-01-02 15:04:05"

func HelperDate2Str(timeIn time.Time) string {
	return timeIn.Format(baseFormat)
}

// HelperStr2Date 时间转字符串
func HelperStr2Date(strTime string) (time.Time, error) {
	return time.Parse(baseFormat, strTime)
}
