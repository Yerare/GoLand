package main

import (
	"easymath/internal/data" // New import
	"fmt"
	"net/http"
	"time" // New import
)

func (app *application) createformulaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new formula")
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
