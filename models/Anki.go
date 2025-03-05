package models

import (
	"math"
	"time"
)

type AnkiValue struct {
	NoteId int
	EF     float64
	Value  int
	Depth  int
	Date   time.Time
	Before int
}

func (note *AnkiValue) GetNewEF(memory int) float64 {
	var EF float64
	if note.Depth == 1 {
		EF = float64(2.5)
	} else {
		EF = float64(note.EF) - float64(0.8) + float64(0.28)*float64(memory) - float64(0.02)*float64(memory*memory)
	}
	return EF
}
func (note *AnkiValue) GetNewI() int {
	if note.Depth == 1 {
		note.Before = 1
		return 1
	}
	if note.Depth == 2 {
		note.Before = 4
		return 4
	}
	note.Before = note.Value
	return int(math.Ceil(float64(note.Before) * note.EF))

}
func (note *AnkiValue) UpdateNewEF(memory int) {
	note.Depth = note.Depth + 1
	//var memory int
	//fmt.Print("记忆情况")
	//fmt.Scanln(&memory)
	NewEF := note.GetNewEF(memory)
	if NewEF < float64(1.3) {
		NewEF = float64(1.3)
	}
	note.EF = NewEF
}
func (note *AnkiValue) Anki(memory int) {
	note.UpdateNewEF(memory)
	note.Value = note.GetNewI()
	//	note.Date =
}

var Memory = make(map[int][]AnkiValue)
