/**
 * @Author: Cc
 * @Description: MysqlClient
 * @File: example
 * @Version: 1.0.0
 * @Date: 2022/8/1 17:37
 * @Software : GoLand
 */

package mysqlc

import (
	"fmt"
	"github.com/Cc360428/HelpPackage/other"
	"gorm.io/gorm"
	"time"
)

var (
	exampleClient *gorm.DB
)

type User struct {
	ID                     int64  `json:"id,omitempty" gorm:"primary_key;comment:'id'"`
	Name                   string `json:"name,omitempty" gorm:"not nul;unique_index:name;UNIQUE;comment:'用户名'"`
	Sex                    int
	Avatar                 string
	Birthday               string
	Email                  string `json:"email,omitempty" gorm:"not nul;unique_index:email;UNIQUE;comment:'邮箱'"` //邮箱
	Password               string
	Salt                   string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAExampleClientt gorm.DeletedAt `gorm:"index"`
}

func InitUserProfile() {

	var count int
	if err := exampleClient.Raw("select count(id) from t_user").Scan(&count).Error; err != nil {
		fmt.Println(err.Error())
	}

	if count == 0 {
		salt, _ := other.Salt()
		exampleClient.Select("Name", "Sex", "Avatar", "Email", "Password", "Salt").Create(&User{
			Name:     "admin",
			Sex:      1,
			Birthday: "1997-06-12",
			Avatar:   "https://gitee.com/hyperli0612/images/raw/master/wy/go/golang.jpeg",
			Email:    "li_chao_cheng@163.com",
			Password: other.Encryption(other.Md5V3("Cc360428"), salt),
			Salt:     salt,
		})
	}
}
