package services

import (
	"Ankipage/db"
	"testing"
)

func TestLoginService(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name  string
		input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
	}{
		{"No.1",
			struct {
				Username string `json:"username" binding:"required"`
				Password string `json:"password" binding:"required"`
			}{
				Username: "用户3",
				Password: "123",
			},
		},
		{"No.2",
			struct {
				Username string `json:"username" binding:"required"`
				Password string `json:"password" binding:"required"`
			}{
				Username: "用户4",
				Password: "123",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			token, ans, err := LoginService(c.input)
			if err != nil {
				t.Errorf("Register err:%v", err)
			} else {
				t.Logf("name:%s\ntoken:%suser%v", c.name, token, &ans)
			}
		})
	}
}
