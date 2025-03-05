package services

import (
	"Ankipage/db"
	"testing"
)

func TestRecentNotes(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name          string
		userid, limit int
	}{
		{"No.1", 10, 4},
		{"No.2", 11, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans, err := RecentNotes(c.userid, c.limit)
			if err != nil {
				t.Errorf("RecentNotes err:%v", err)
			} else {
				t.Logf("name:%s\nRecentNotes:%v", c.name, ans)
			}
		})
	}
}
