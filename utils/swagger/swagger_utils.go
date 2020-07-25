package swagger

import (
	"encoding/json"
	"github.com/Cc360428/HelpPackage/utils/logs"
	"io/ioutil"
	"time"
)

// 获取swagger.json 内容
func GetSwagger(path string) map[string]interface{} {
	// 读取文件
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	// 将[]byte转为 map
	allJson2Map := make(map[string]interface{})
	err = json.Unmarshal(file, &allJson2Map)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	// 获取转为 map 中的 "paths"
	f := allJson2Map["paths"].(map[string]interface{})
	return f
}

type Swagger struct {
	Request
	Url    string `json:"url"`    // url
	Method string `json:"method"` // 方法
}

type Request struct {
	Tags    []string `json:"tags"`    // 属于那个模块
	Summary string   `json:"summary"` // 方法名称
}

// 获取swagger 解析出来的在解析具体内容
func GetDetails(urlDetails interface{}) Swagger {
	// 将 传进来的interface{} 类型转为 map 对象
	z := make(map[string]interface{})
	err := InterfaceToInterface(urlDetails, &z)
	if err != nil {
		logs.Error(err.Error())
	}
	var swagger Swagger
	for k, v := range z {
		swagger.Method = k
		var requestDetails Request
		marshal, err := json.Marshal(v)
		if err != nil {
			logs.Error(err.Error())
		}
		err = json.Unmarshal(marshal, &requestDetails)
		if err != nil {
			logs.Error(err.Error())
		}
		swagger.Request = requestDetails
	}
	return swagger
}

// 不同类型转换
func InterfaceToInterface(in, out interface{}) error {
	marshal, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, out)
}

// 启动初始化添加到DB
func init() {
	var requestAll []Swagger
	for key, value := range GetSwagger("swagger/swagger.json") {
		details := GetDetails(value)
		details.Url = key
		requestAll = append(requestAll, details)
	}
	var add Resources
	for _, value := range requestAll {
		value := value
		go func() {
			add.Method = value.Method
			add.Summary = value.Summary
			add.URL = value.Url
			add.Tags = value.Tags[0]
			err := AddResource(add)
			if err != nil {
				panic(err.Error())
			}
		}()
	}
}

// 返回结构
type Resources struct {
	ID        int64      `json:"id,omitempty" gorm:"primary_key comment:'ID'"`
	Tags      string     `json:"tags" gorm:"comment:'属于那个模块'"`             // 属于那个模块
	Summary   string     `json:"summary" gorm:"comment:'方法名称'"`            // 方法名称
	URL       string     `json:"url" gorm:"not null;unique;comment:'url'"` // url
	Method    string     `json:"method" gorm:"comment:'请求方法'"`             // 请求方法
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"comment:'更新时间'"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"  sql:"index" gorm:"comment:'删除时间软删除'"`
}

// 添加至数据库
func AddResource(r Resources) error {
	//sql := `insert into cc_resources (tags, summary,url,method,created_at,updated_at) values(?,?,?,?,NOW(),NOW()) ON DUPLICATE KEY UPDATE tags=?, summary=?,url=?,method=?,updated_at=NOW()`
	//err := MysqlClient.Exec(sql, r.Tags, r.Summary, r.URL, r.Method, r.Tags, r.Summary, r.URL, r.Method).Error
	//if err != nil {
	//	logs.Error(err.Error())
	//	return err
	//}
	return nil
}
