package services

import (
	"Ankipage/db"
	"testing"
)

func TestDeleteNote(t *testing.T) {
	db.InitDB()
	cases := []struct {
		name string
		id   int
	}{
		{"No.1", 10},
		{"No.2", 11},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := DeleteNote(c.id)
			if err != nil {
				t.Errorf("DeleteNote err:%v", err)
			} else {
				t.Log("DeleteNote success")
			}
		})
	}
}
