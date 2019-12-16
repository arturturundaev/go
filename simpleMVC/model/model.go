package model

import (
	"fmt"
	uuid "github.com/google/uuid"
)

type User struct {
	id 		 string
	login 	 string
	password string
}

var EntityArr = map[string]User{}


func Get(id string) User {
	return EntityArr[id]
}


func GetAll() map[string]User {
	uuidString := uuid.New().String()
	user := User{id:uuidString, login:"myLogin", password:"MyPassword"}
	EntityArr[uuidString] = user

	return EntityArr;
}


func Save(id string, data map[string]string) {

	currentUser := EntityArr[id]
	currentUser.login = data["login"]
	currentUser.password = data["password"]

	EntityArr[id] = currentUser
}


func Delete(id string) {
	delete(EntityArr, id)
	fmt.Println(EntityArr)
}