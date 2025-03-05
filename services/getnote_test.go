package services

import (
	"Ankipage/db"
	"testing"
)

func TestGetNote(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name   string
		userid int
	}{
		{"No.1", 10},
		{"No.2", 11},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans, err := GetNote(c.userid)
			if err != nil {
				t.Errorf("GetNote err:%v", err)
			} else {
				t.Logf("name:%s\nGetNote:%v", c.name, ans)
			}
		})
	}
}
