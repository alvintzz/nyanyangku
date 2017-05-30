package middleware

import (
	"net/http"
    "log"

	"github.com/alvintzz/nyanyangku/common/render"
)

type HandlerLayout func(rw http.ResponseWriter, r *http.Request) (string, map[string]interface{}, error)
func (fn HandlerLayout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	include, param, _ := fn(w, r)

	param["include"] = include

    e, _ := render.Get("main")
	err := e.Template.ExecuteTemplate(w, "layout", param)
	if err != nil {
		log.Println(err)
	}
}

type HandlerTemplate func(rw http.ResponseWriter, r *http.Request) (string, map[string]interface{}, error)
func (fn HandlerTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	include, param, _ := fn(w, r)

    e, _ := render.Get("main")
	err := e.Template.ExecuteTemplate(w, include, param)
	if err != nil {
		log.Println(err)
	}
}
