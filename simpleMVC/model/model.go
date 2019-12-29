package model

import (
	"fmt"
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
	return EntityArr
}

func Update(data map[string]string) {
	fmt.Println("Update")				
	user := getEntity(data)
	EntityArr[user.Id] = user
}

func Create(data map[string]string) bool {
	fmt.Println("Create")
	user := getEntity(data)
	fmt.Println("Create")				
	fmt.Println(user)				
	EntityArr[user.Id] = user

	return true
}

func Delete(id string) {
	delete(EntityArr, id)
}


func getEntity(data map[string]string)Entity {
	fmt.Println("GetEntity")
	fmt.Println(data)
	if(data["Id"] == "") {
		data["Id"] = uuid.New().String()
	}

	return Entity{	Id: 	  data["Id"], 
					Login: 	  data["Login"], 
					Password: data["Password"]}
}