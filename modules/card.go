package modules

import (
	"time"
)

type Card struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	TimeGuessForDone  int        `json:"timeGuessForDone"`
	TimeRealForDone   int        `json:"timeRealForDone"`
	Due               *time.Time `json:"due"`
	DateLastChangeDue *time.Time `json:"dateLastActivity"`
}
