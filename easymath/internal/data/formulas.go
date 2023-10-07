package data

import (
	"time"
)

type Formulas struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Chapter   string    `json:"chapter"`
	Level     string    `json:"levels,omitempty"`
	Withvar   Withvar   `json:"count"`
}
