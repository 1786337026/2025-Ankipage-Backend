package services

import (
	"Ankipage/db"
	"testing"
)

func TestUpdateNote(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name    string
		id      int
		content string
	}{
		{"No.1", 10, "修改后的No.1笔记"},
		{"No.2", 11, "修改后的No.2笔记"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := UpdateNote(c.id, c.content)
			if err != nil {
				t.Errorf("UpdateNote err:%v", err)
			} else {
				t.Logf("name:%s\ncontent:%s", c.name, c.content)
			}
		})
	}
}
