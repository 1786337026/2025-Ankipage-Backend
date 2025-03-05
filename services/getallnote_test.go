package services

import (
	"Ankipage/db"
	"testing"
)

func TestGetAllNote(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name   string
		userid int
	}{
		{"No.1", 1},
		{"No.2", 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans, err := GetAllNote(c.userid)
			if err != nil {
				t.Errorf("GetAllNote err:%v", err)
			} else {
				t.Logf("name:%s\nGetAllNote:%v", c.name, ans)
			}
		})
	}
}
