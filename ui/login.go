package ui

import (
	"net/http"
	"fmt"
	"log"

	"github.com/alvintzz/nyanyangku/model"
	"github.com/alvintzz/nyanyangku/common/database"
)

func LoginFormHandler(w http.ResponseWriter, r *http.Request) (string, map[string]interface{}, error) {
	context := map[string]interface{}{}
	return "login", context, nil
	// engine.RenderPlain(w, "login", context)
}


type UserLogin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
func LoginActionAjaxHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	userEmail := r.FormValue("input_email")
	userPassword := r.FormValue("input_password")
	login := UserLogin{}

	masterDB, _ := database.Get("main")
	userModel := model.NewUserModel(masterDB)
	user, err := userModel.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err)
		return login, err
	}
	if user.ID == 0 {
		return login, fmt.Errorf("User Not found")
	}
	if user.Password != userPassword {
		return login, fmt.Errorf("Wrong Password")
	}

	login.UserEmail = userEmail
	login.UserPassword = userPassword
	return login, nil
}
