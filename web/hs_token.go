// Package web token
package web

import (
	"errors"
	"github.com/Cc360428/HelpPackage/other"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	UserType int64  `json:"user_type"`
	RegionId int64  `json:"region_id"`
}

// CreateToken 创建token
func CreateToken(user *UserInfo) (tokens string, err error) {
	//自定义claim
	claim := jwt.MapClaims{
		"id":        user.Id,
		"user_name": user.UserName,
		"region_id": user.RegionId,
		"user_type": user.UserType,
		"nbf":       time.Now().Unix(),
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 480).Unix(), //有效期20天
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err = token.SignedString([]byte(token.Signature))
	return
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(token.Signature), nil
	}
}

// ParseToken 获取token中的结构体
func ParseToken(tokens string) (user *UserInfo, err error) {
	user = &UserInfo{}
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	user.Id, err = other.ToInt64(claim["id"])
	user.UserType, _ = other.ToInt64(claim["user_type"])
	user.RegionId, _ = other.ToInt64(claim["region_id"])
	user.UserName = claim["user_name"].(string)
	return
}
