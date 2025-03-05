package services

import (
	"testing"
)

func TestAnki(t *testing.T) {
	cases := []struct {
		name   string
		userid int
	}{
		{"No.1", 1},
		{"No.2", 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			notes, err, _ := Anki(c.userid)
			//assert.Equal(t, nil, err)
			if err != nil {
				t.Errorf("Anki error = %v", err)
			} else {
				t.Logf("name = %s\nnotes = %v", c.name, notes)
			}
		})
	}
}
