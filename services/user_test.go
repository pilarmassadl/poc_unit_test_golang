package services_test

import (
	"BasicTestGuild/model"
	"BasicTestGuild/services"
	"BasicTestGuild/util"
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestUserService_GetAll(t *testing.T) {

	util.UserList = append(util.UserList, model.User{
		Document: "1111",
		Name:     "Test",
		Email:    "test",
		Phone:    "23456",
		Age:      0,
	})

	userService := services.NewUser()

	userList, _:= userService.GetAll()

	if len(*userList) == 0 {
		t.Error("No hay almacenados")
		t.Fail()
	}else {
		t.Log("Test OK")
	}

}


func TestUserService_Create(t *testing.T) {

	userService := services.NewUser()

	var body bytes.Buffer
	body.WriteString("{\"document\":\"1\",\"name\":\"Pili2\",\"email\":\"pilarica@gmail.com\",\"phone\":\"3138320434\",\"age\":27}")

	req:= httptest.NewRequest("POST","http://localhost:8084/user" ,&body)

	newUser, _ :=userService.Create(req)

	if newUser == nil {
		t.Error("Error al crear usuario")
		t.Fail()
	} else if newUser.Name!="Pili"{
		t.Errorf("Nombre de usuario no corresponde el esperado es:%s ,actual: %s", "Pili", newUser.Name  )
		t.Fail()
	} else {
		t.Log("Test OK")
	}

}