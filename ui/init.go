package ui

import (
	"net/http"

	"github.com/alvintzz/nyanyangku/util/middleware"
)

func InitHandler(mux *http.ServeMux, templateLocation string) {
	mux.Handle("/login", middleware.HandlerTemplate(LoginFormHandler))
	mux.Handle("/ajax/login", middleware.HandlerJSON(LoginActionAjaxHandler))
}
