/**
 * @Author: Cc
 * @Description: Apollo 配置中心
 * @File: apollo
 * @Version: 1.0.0
 * @Date: 2022/12/15 14:32
 * @Software : GoLand
 */

package apollo

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var osType string
var PathSeparator string

const WINDOWS = "windows"

type ConfigFile struct {
	NameSpace      string //apollo 对应的namespaces(文件名)
	Path           string //配置路径
	notificationID int
}

type Notification struct {
	NamespaceName  string `json:"namespaceName"`
	NotificationID int    `json:"notificationId"`
}

type ApolloClient struct {
	ClusterName string
	Server      string
	AppID       string
	ReleaseKey  string
	IP          string
	done        chan struct{}
	configFiles sync.Map
	client      *http.Client
}

func init() {
	osType = runtime.GOOS
	if osType == WINDOWS { //前边的判断是否是系统的分隔符
		PathSeparator = "\\"
	} else {
		PathSeparator = "/"
	}
}

func NewApolloClient(address string, clusterName string, appId string) *ApolloClient {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				//内网使用，不强制校验服务端的证书
				InsecureSkipVerify: true,
			},
			DisableKeepAlives:   false,
			MaxIdleConns:        20,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     60 * time.Second,
		},
		Timeout: time.Second * 90,
	}

	return &ApolloClient{
		ClusterName: clusterName,
		Server:      address,
		AppID:       appId,
		ReleaseKey:  "",
		IP:          "",
		done:        make(chan struct{}, 1),
		configFiles: sync.Map{},
		client:      client,
	}
}

func (p *ApolloClient) Start() (err error) {
	err = p.start()
	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case <-time.After(time.Second * 120):
				p.start()
			case <-p.done:
				log.Println("apollo listen done")
				return
			}
		}
	}()
	return nil
}

func (p *ApolloClient) start() (err error) {
	notifyFiles, err := p.getNotification()
	if err != nil {
		log.Println("failed to get notification,err:", err)
		return
	}

	for nameSpaceName, notifycationId := range notifyFiles {
		content, err1 := p.GetConfig(nameSpaceName)
		if err1 != nil {
			log.Println("failed to get config form apollo,err:", err1, "namespace:", nameSpaceName)
			return err1
		}

		configFile := p.getConfigFile(nameSpaceName)
		if configFile == nil {
			log.Println("failed to get config form apollo,err:", err1, "namespace:", nameSpaceName)
			return err1
		}

		err1 = p.reWrite(*configFile, []byte(content))
		if err1 != nil {
			log.Println("failed to rewrite config", configFile, "err:", err1)
			return err1
		}

		//重置notificationId
		configFile.notificationID = notifycationId
		p.configFiles.Store(nameSpaceName, configFile)
	}

	return
}

func (p *ApolloClient) Stop() {
	p.done <- struct{}{}
}

// 查询被修改的文件
func (p *ApolloClient) getNotification() (notifyFiles map[string]int, err error) {

	notifyFiles = make(map[string]int)

	// URL: {config_server_url}/notifications/v2?appId={appId}&cluster={clusterName}&notifications={notifications}
	//Method: GET
	notifications := make([]Notification, 0)

	params := url.Values{}
	params.Add("appId", p.AppID)
	params.Add("cluster", p.ClusterName)

	p.configFiles.Range(func(key, value interface{}) bool {
		v, ok := value.(*ConfigFile)
		if !ok {
			return true
		}

		notifications = append(notifications, Notification{
			NamespaceName:  v.NameSpace,
			NotificationID: v.notificationID,
		})
		return true
	})

	body, err := json.Marshal(notifications)
	if err != nil {
		log.Println("failed to marshal notifications,err:", err)
		return
	}

	params.Add("notifications", string(body))

	uri := fmt.Sprintf(
		"%s/notifications/v2?%s",
		p.Server,
		params.Encode(),
	)

	code, res, err := p.request(http.MethodGet, uri, "", nil)
	if err != nil {
		log.Println("从配置中心拉取配置修改信息失败，err:", err, "url:", uri, "code:", code)
		return
	}

	//如果返回没有变化，则直接返回
	if code == http.StatusNotModified {
		return
	}

	//不等于成功，报错
	if code != http.StatusOK {
		log.Println("从配置中心拉取配置修改信息失败，err:", err, "url:", uri, "code:", code)
		return
	}

	//解析出来，有被修改的namespace(文件)
	// Notification apollo的通知结构
	var result []Notification

	err = json.Unmarshal(res, &result)
	if err != nil {
		log.Println("failed to unmarshal notification,err:", err)
		return
	}

	for _, v := range result {
		notifyFiles[v.NamespaceName] = v.NotificationID
		log.Println("notification:", v.NamespaceName, v.NotificationID)
	}
	return
}

func (p *ApolloClient) GetConfig(nameSpace string) (content string, err error) {

	//URL: {config_server_url}/configs/{appId}/{clusterName}/{namespaceName}?releaseKey={releaseKey}&ip={clientIp}
	//Method: GET
	uri := fmt.Sprintf(
		"%s/configs/%s/%s/%s",
		p.Server,
		p.AppID,
		p.ClusterName,
		nameSpace,
	)

	params := url.Values{}
	if p.IP != "" {
		params.Add("ip", p.IP)
		uri = uri + "?" + params.Encode()
	}

	var result struct {
		Configurations map[string]string `json:"configurations"`
		ReleaseKey     string            `json:"releaseKey"`
		Cluster        string            `json:"cluster"`
		NamespaceName  string            `json:"namespaceName"`
		AppID          string            `json:"appId"`
	}

	code, res, err := p.request(http.MethodGet, uri, "", nil)
	if err != nil || code != http.StatusOK {
		log.Println("从配置中心拉取配置失败，err:", err, "url:", uri, "code:", code)
		return
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		log.Println("failed to unmarshal body,err:", err)
		return
	}

	return result.Configurations["content"], err
}

// 增加监听配置文件修改
func (p *ApolloClient) AddConfig(configFiles []ConfigFile) (err error) {

	for _, v := range configFiles {
		configFile := v
		key := configFile.NameSpace
		configFile.notificationID = -1
		_, ok := p.configFiles.Load(key)
		if ok {
			log.Println("already listen config:", key)
			return
		}

		_, err = p.GetConfig(configFile.NameSpace)
		if err != nil {
			log.Println("failed to get config from apollo,err:", err)
			return
		}
		p.configFiles.Store(key, &configFile)
	}

	return
}

// 写文件
func (p *ApolloClient) reWrite(configFile ConfigFile, data []byte) (err error) {

	exist, err := pathExists(configFile.Path)
	if !exist || err != nil {
		err = os.Mkdir(configFile.Path, 0644)
		if err != nil {
			return
		}
	}

	//空内容，不覆盖本地
	if len(data) == 0 {
		log.Println("content is empty,config file:", configFile)
		return
	}

	//TODO 是否需要更安全的校验？
	fileName := strings.Join([]string{configFile.Path, configFile.NameSpace}, PathSeparator)
	log.Println("rewrite file:", fileName)
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Println("failed to write file:", configFile.Path, "err:", err)
		return
	}
	return
}

func (p *ApolloClient) getConfigFile(nameSpaceName string) *ConfigFile {
	v, ok := p.configFiles.Load(nameSpaceName)
	if !ok {
		log.Println("not exist namespace name:", nameSpaceName)
		return &ConfigFile{}
	}

	configFile, ok := v.(*ConfigFile)
	if !ok {
		log.Println("not exist namespace name:", nameSpaceName)
		return &ConfigFile{}
	}

	return configFile
}

func (p *ApolloClient) request(method, url, body string, header map[string]string) (code int, res []byte, err error) {
	//判断请求的方法是否正确
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
	default:
		fmt.Println("[ERROR] invalid request method:", method)
		err = fmt.Errorf("invalid request method%v", method)
		return
	}

	//判断url是否有效
	if "" == url {
		fmt.Println("[ERROR] request url is empty")
		err = fmt.Errorf("request url is empty")
		return
	}

	strings.NewReader(body)
	//创建一个request
	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		fmt.Println("[ERROR] internal server error:", err)
		return
	}

	if len(header) > 0 {
		for k, v := range header {
			if "" != k {
				r.Header.Set(k, v)
			}
		}
	}

	//设置代理信息
	r.Header.Set("User-Agent", "huoys")

	resp, err := p.client.Do(r)
	if err != nil {
		fmt.Println("[ERROR] failed to do respond ,err:", err)
		return
	}

	//限制响应消息体,防止获得恶意攻击信息
	bodyRead := http.MaxBytesReader(nil, resp.Body, 1024*10240)

	code = resp.StatusCode
	defer bodyRead.Close()
	res, err = ioutil.ReadAll(bodyRead)
	if err != nil {
		fmt.Println("[ERROR] failed to get respond body,err:", err)
		return
	}

	return
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
