package main

import (
	"easymath/internal/data" // New import
	"easymath/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createFormulasHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Chapter string       `json:"Chapter"`
		Level   string       `json:"Level"`
		Withvar data.Withvar `json:"Withvar"`
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
	v := validator.New()
	if data.ValidateFormulas(v, formulas); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/formulas/%d", formulas.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"formulas": formulas}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	v.Check(input.Chapter != "", "chapter", "must be provided")
	v.Check(len(input.Chapter) <= 500, "chapter", "must not be more than 500 bytes long")
	v.Check(input.Withvar != 0, "count", "must be provided")
	v.Check(input.Withvar > 0, "count", "must be a positive integer")
	v.Check(input.Level != "", "levels", "must be provided")
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

func (app *application) showFormulasHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	formulas, err := app.models.Formulas.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"formulas": formulas}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) updateFormulasHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	formula, err := app.models.Formulas.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Chapter string       `json:"Chapter"`
		Level   string       `json:"Levels"`
		Withvar data.Withvar `json:"Withvar"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Chapter != nil {
		formulas.Chapter = *input.Chapter
	}
	if input.Level != nil {
		formulas.Level = *input.Levels
	}
	if input.Withvar != nil {
		formulas.Withvar = *input.Withvar
	}
	v := validator.New()
	if data.ValidateFormulas(v, formula); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Formulas.Update(formula)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"formula": formula}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteFormulasHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Formulas.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Formula successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) listFormulasHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Chapter string
		Level   string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()
	input.Chapter = app.readString(qs, "Chapter", "")
	input.Level = app.readString(qs, "Level", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "chapter", "level", "withvar", "-id", "-chapter", "-level", "-withvar"}
	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	formulas, metadata, err := app.models.Formulas.GetAll(input.Chapter, input.Level, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"formulas": formulas, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
