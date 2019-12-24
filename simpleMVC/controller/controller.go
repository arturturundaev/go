package controller



import (
	"fmt"

	user "../model"
)

type UserController struct {

}

func (cntr UserController)ShowAll() map[string]user.Entity {
	return user.GetAll()
}



func Add(formData map[string]string) user.Entity {
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

func Get(id string)user.Entity {
	return user.Get(id)
}
