package controller



import (
	"fmt"
	user "../model"
	"plugin"
)

type UserController struct {

}

func (cntr UserController)ShowAll() map[string]user.User  {
	return user.GetAll()
}



func Add(formData map[string]string) user.User {
  //if formData["Id"] != "" {
  	return user.Update(formData["Id"], formData["Login"], formData["Password"])
  //}
}


func Delete(id string) {
  user.Delete(id)
}

func Edit() {
  fmt.Println("EDIT")
}

func Get(id string)user.User {
	return user.Get(id)
}
