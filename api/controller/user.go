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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserCRUD(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindAll()
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		response.JSON(w, http.StatusOK, users)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserCRUD(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.Save(user)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		response.JSON(w, http.StatusCreated, user)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryUserCRUD(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindById(uint(userId))
		if err != nil {
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}
		response.JSON(w, http.StatusOK, user)
	}(repo)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserCRUD(db)

	func(userRepository repository.UserRepository) {
		rows, err := userRepository.Update(uint(userId), user)
		if err != nil {
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}
		response.JSON(w, http.StatusOK, rows)
	}(repo)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
