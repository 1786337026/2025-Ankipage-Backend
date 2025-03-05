package services

import (
	"Ankipage/db"
	"testing"
)

func TestRegister(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name                      string
		username, password, email string
	}{
		{"No.1", "用户3", "123", "email3"},
		{"No.2", "用户4", "123", "email4"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans, err := Register(c.username, c.password, c.email)
			if err != nil {
				t.Errorf("Register err:%v", err)
			} else {
				t.Logf("name:%s\nuser:%v", c.name, ans)
			}
		})
	}
}
