package controller

import (
	"fmt"
	user "../model"
)


func ShowAll() map[string]user.User  {
	return user.GetAll()
}



func Add() {
  fmt.Println("DELETE")
}


func Delete(id string) {
  user.Delete(id)
}

func Edit() {
  fmt.Println("EDIT")
}

func Get() {
   fmt.Println("GET")
}
