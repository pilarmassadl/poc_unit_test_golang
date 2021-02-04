package services

import (
	"BasicTestGuild/model"
	"BasicTestGuild/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewUser() UserService {
	return &userService{}
}

type userService struct {
}

type UserService interface {
	Create(*http.Request) (*model.User,error)
	GetAll() (*model.UserList,error)
}

func (s *userService) Create(r *http.Request) (user *model.User,err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Sprint("Error tratando de decodificar el usuario")
		return
	}
	util.UserList = append(util.UserList, *user)
	return
}

func (s *userService) GetAll() (list *model.UserList,err error) {
	return &util.UserList, nil
}