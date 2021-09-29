package jwt

import (
	"testing"
	"time"
)

type User struct {
	Id       int
	UserName string
}

func TestNewAuthToken(t *testing.T) {
	users := &User{Id: 1, UserName: "wobusbzzzzzzzzzzzz"}
	authToken := NewAuthToken()
	token, err := authToken.EncodeToken(users)

	if err != nil {
		t.Errorf("create token failed error %s\n", err)
	}
	t.Log(token)
	var (
		curTime = time.Unix(time.Now().Unix(), 0)
	)

	if autoToken, err := NewAuthToken().DecodeToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODc0MDA1MjgsImlhdCI6MTU4Njc5NTcyOCwiaXNzIjoid29idXNienMiLCJuYmYiOjE1ODY3OTU3Mjh9.c5HK3r6V98SY3RSe2JUuCBio3opUGii0MzpqKkI38tQ"); err == nil {
		if user, ok := autoToken.UserInter.(*User); !ok {
			t.Errorf("NotBefore: %v  want = %v\n", user, users)
		}
		if curTime.Hour() != time.Unix(autoToken.IssuedAt, 0).Hour() {
			t.Errorf("NotBefore: %s  want = %d\n", time.Unix(autoToken.NotBefore, 0).Format("2006-01-02 15:04:05"), curTime.Day())
		}
		if curTime.Hour() != time.Unix(autoToken.IssuedAt, 0).Hour() {
			t.Errorf("IssuedAt: %s  want = %d\n", time.Unix(autoToken.IssuedAt, 0).Format("2006-01-02 15:04:05"), curTime.Day())
		}
		if time.Unix(autoToken.ExpiresAt, 0).Day() != curTime.Day()+7 {
			t.Errorf("ExpiresAt: %s want = %s\n", time.Unix(autoToken.ExpiresAt, 0).Format("2006-01-02 15:04:05"), curTime.Format("2006-01-02 15:04:05"))
		}
	} else {
		t.Errorf("decode token failed error %s", err)
	}
}
