package model

import (
	uuid "github.com/google/uuid"
)

type Entity struct {
	Id       string
	Login    string
	Password string
}

var EntityArr = make(map[string]interface{})

func Get(id string) interface{} {
	return EntityArr[id]
}

func GetAll() map[string]interface{} {
	uuidString := uuid.New().String()
	user := Entity{Id: uuidString, Login: "myLogin", Password: "MyPassword"}
	EntityArr[uuidString] = user

	return EntityArr
}

func Update(id string, login string, password string) interface{} {

	currentUser := EntityArr[id]
	//currentUserLogin = login
	//currentUserPassword = password

	EntityArr[id] = currentUser

	return currentUser
}

func Delete(id string) {
	delete(EntityArr, id)
}
