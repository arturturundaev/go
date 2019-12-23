package model

import (
	"fmt"
	uuid "github.com/google/uuid"
)

type User struct {
	Id       string
	Login    string
	Password string
}

var EntityArr = map[string]User{}


func Get(id string) User {
	return EntityArr[id]
}


func GetAll() map[string]User {
	uuidString := uuid.New().String()
	user := User{Id: uuidString, Login:"myLogin", Password:"MyPassword"}
	EntityArr[uuidString] = user

	return EntityArr
}


func Update(id string, login string, password string) User {

	currentUser := EntityArr[id]
	currentUser.Login = login
	currentUser.Password = password

	EntityArr[id] = currentUser

	return currentUser
}


func Delete(id string) {
	delete(EntityArr, id)
	fmt.Println(EntityArr)
}