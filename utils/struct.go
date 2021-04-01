package utils

import (
	"fmt"
	"strconv"
)

type RedisUser struct {
	Id     int64  `json:"id"`
	PropId string `json:"prop_id"`
}

func RedisGo() []RedisUser {
	var user []RedisUser
	var i int64
	for i = 1; i < 10; i++ {
		user = append(user, RedisUser{
			Id:     i,
			PropId: strconv.FormatInt(i, 10) + "lcc-redis-ligo",
		})
	}
	return user
}

type HallUserInfo struct {
	Id   int64  `json:"id"`
	Nick string `json:"nick"`
	Age  int    `json:"age"`
}

func HallUserInfoGo() []HallUserInfo {
	var user []HallUserInfo
	user = append(user, HallUserInfo{
		Id:   2,
		Nick: "2-GetHall",
		Age:  23,
	})
	user = append(user, HallUserInfo{
		Id:   3,
		Nick: "3-GetHall",
		Age:  33,
	})
	user = append(user, HallUserInfo{
		Id:   6,
		Nick: "6-GetHall",
		Age:  66,
	})
	return user
}

type UserAll struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	PropId   string `json:"prop_id"`
	Nick     string `json:"nick"`
	Age      int    `json:"age"`
}

func TestHall() {
	redisGo := RedisGo()
	infoGo := HallUserInfoGo()
	mapHallUserInfo := make(map[int64]HallUserInfo)
	for _, info := range infoGo {
		mapHallUserInfo[info.Id] = info
	}
	userAll := make([]UserAll, len(redisGo))
	for index, value := range redisGo {
		userAll[index].PropId = value.PropId
		userAll[index].Id = value.Id
		userAll[index].Age = mapHallUserInfo[value.Id].Age
		userAll[index].Nick = mapHallUserInfo[value.Id].Nick
	}
	fmt.Println(userAll)
}
