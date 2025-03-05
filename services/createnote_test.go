package services

import (
	"Ankipage/db"
	"testing"
)

func TestCreateNote(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name       string
		id, userid int
		content    string
	}{
		{"No.1", 10, 1, "这是第一个测试笔记"},
		{"No.2", 11, 1, "这是第二个测试笔记"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := CreateNote(c.id, c.userid, c.content)
			//assert.Equal(t, err)
			if err != nil {
				t.Errorf("CreateNote err:%v", err)
			} else {
				t.Logf("name:%s\ncontent:%s", c.name, c.content)
			}
		})
	}
}
