package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/termins", app.terminsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/formulas", app.requirePermission("formulas:read", app.listFormulasHandler))
	router.HandlerFunc(http.MethodPost, "/v1/formulas", app.requirePermission("formulas:write", app.createFormulasHandler))
	router.HandlerFunc(http.MethodGet, "/v1/formulas/:id", app.requirePermission("formulas:read", app.showFormulasHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/formulas/:id", app.requirePermission("formulas:write", app.updateFormulasHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/formulas/:id", app.requirePermission("formulas:write", app.deleteFormulasHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
