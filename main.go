package main

import (
	"BasicTestGuild/services"
	"encoding/json"
	"net/http"
	"strings"
)

func CrudUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		if strings.HasSuffix(r.URL.Path, "/user"){
			getAll(w)
		}

	case "POST":
		save(w,r)
	default:
		getAll(w)
	}

}

func getAll(w http.ResponseWriter){
	userService:= services.NewUser()
	userList,err:= userService.GetAll()

	js, err := json.Marshal(userList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}

func save(w http.ResponseWriter, r *http.Request){
	userService:= services.NewUser()

	newUser,err:= userService.Create(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(js)
}

func main() {
	http.HandleFunc("/user", CrudUser)
	http.HandleFunc("/user/", CrudUser)
	http.ListenAndServe(":8084", nil)
}