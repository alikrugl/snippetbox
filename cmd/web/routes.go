package main

import (
	"net/http"

	"github.com/alikrugl/snippetbox/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))

	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))

	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))
	mux.Handle("POST /user/logout", dynamic.ThenFunc(app.userLogoutPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /snippet/view/{id}", protected.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
