package controller

import "net/http"

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("listing all users"))
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get a user"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
