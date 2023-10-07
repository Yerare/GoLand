package main

import (
	"easymath/internal/data" // New import
	"easymath/internal/validator"
	"fmt"
	"net/http"
	"time" // New import
)

func (app *application) createformulaHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Chapter string       `json:"chapter"`
		Level   string       `json:"levels"`
		Withvar data.Withvar `json:"count"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	formulas := &data.Formulas{
		Chapter: input.Chapter,
		Withvar: input.Withvar,
		Level:   input.Level,
	}
	// Initialize a new Validator instance.
	v := validator.New()
	// Use the Check() method to execute our validation checks. This will add the
	// provided key and error message to the errors map if the check does not evaluate
	// to true. For example, in the first line here we "check that the title is not
	// equal to the empty string". In the second, we "check that the length of the title
	// is less than or equal to 500 bytes" and so on.
	v.Check(input.Chapter != "", "chapter", "must be provided")
	v.Check(len(input.Chapter) <= 500, "chapter", "must not be more than 500 bytes long")
	v.Check(input.Withvar != 0, "runtime", "must be provided")
	v.Check(input.Withvar > 0, "runtime", "must be a positive integer")
	v.Check(input.Level != "", "chapter", "must be provided")
	v.Check(len(input.Level) <= 500, "levels", "must not be more than 500 bytes long")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
	if data.ValidateFormulas(v, formulas); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showformulaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	formulas := data.Formulas{
		ID:        id,
		CreatedAt: time.Now(),
		Chapter:   "circle",
		Level:     "easy",
		Withvar:   2,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"formula": formulas}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
