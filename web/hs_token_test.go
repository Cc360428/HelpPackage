package web

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(&UserInfo{
		Id:       1,
		UserName: "Cc",
		UserType: 1,
		RegionId: 1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("this is token:", token)
}

func TestParseToken(t *testing.T) {
	userInfo, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwODYwNTQsImlhdCI6MTY4MzM1ODA1NCwiaWQiOjEsIm5iZiI6MTY4MzM1ODA1NCwicmVnaW9uX2lkIjoxLCJ1c2VyX25hbWUiOiJDYyIsInVzZXJfdHlwZSI6MX0.EttxbdI5-Rl5Myx14cGFixweopUSFYDjTbJ-KfDbWOU")
	if err != nil {
		t.Error(err)
	}
	t.Log("this is UserInfo:", userInfo)

}
