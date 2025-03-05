package services

import "testing"

func TestAnkiPush(t *testing.T) {
	cases := []struct {
		name               string
		id, memory, userid int //memory 0~5
	}{
		//{"No.1", 1},
		//{"No.2", 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := AnkiPush(c.id, c.memory, c.userid)
			if err != nil {
				t.Errorf("Anki error = %v", err)
			} else {
				t.Logf("push successfuly")
			}
		})
	}
}
