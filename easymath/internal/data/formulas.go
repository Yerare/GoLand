package data

import (
	"easymath/internal/validator"
	"time"
)

type Formulas struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Chapter   string    `json:"chapter"`
	Level     string    `json:"levels,omitempty"`
	Withvar   Withvar   `json:"count"`
}

func ValidateFormulas(v *validator.Validator, formula *Formulas) {
	v.Check(formula.Chapter != "", "chapter", "must be provided")
	v.Check(len(formula.Chapter) <= 500, "chapter", "must not be more than 500 bytes long")
	v.Check(formula.Withvar != 0, "count", "must be provided")
	v.Check(formula.Withvar > 0, "count", "must be a positive integer")
	v.Check(formula.Level != "", "levels", "must be provided")
	v.Check(len(formula.Level) <= 500, "levels", "must not be more than 500 bytes long")

}
