package ui

import (
	"net/http"

	"github.com/alvintzz/nyanyangku/util/middleware"
	"github.com/alvintzz/nyanyangku/common/database"
)

var masterDB *database.Db

func InitDb(db *database.Db) {
	masterDB = db
}

func InitHandler(mux *http.ServeMux, templateLocation string) {
	mux.Handle("/login", middleware.HandlerTemplate(LoginFormHandler))
	mux.Handle("/ajax/login", middleware.HandlerJSON(LoginActionAjaxHandler))
}
