package controller

import (
	"encoding/json"
	"github.com/spootrick/survi/api/auth"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/response"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Verify(model.Login)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusOK, token)
}
