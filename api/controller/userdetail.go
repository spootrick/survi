package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spootrick/survi/api/database"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/response"
	"github.com/spootrick/survi/repository"
	"github.com/spootrick/survi/repository/crud"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUserDetail(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	userDetail := model.UserDetail{}
	err = json.Unmarshal(body, &userDetail)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userDetail.Prepare()
	err = userDetail.Verify()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserDetailsCRUD(db)

	func(userDetailRepository repository.UserDetailRepository) {
		user, err := userDetailRepository.Save(userDetail)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		response.JSON(w, http.StatusCreated, user)
	}(repo)
}

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserDetailsCRUD(db)

	func(userDetailRepository repository.UserDetailRepository) {
		userDetail, err := userDetailRepository.FindById(uint(userId))
		if err != nil {
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}
		response.JSON(w, http.StatusOK, userDetail)
	}(repo)
}

func UpdateUserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userDetail := model.UserDetail{}
	err = json.Unmarshal(body, &userDetail)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserDetailsCRUD(db)

	func(userDetailsRepository repository.UserDetailRepository) {
		rows, err := userDetailsRepository.Update(uint(userId), userDetail)
		if err != nil {
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}
		response.JSON(w, http.StatusOK, rows)
	}(repo)
}
