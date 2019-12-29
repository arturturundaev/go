package userController



import (
	user "../../model"
)


type UserController struct {

}


func (cntr UserController)ShowAll() map[string]interface{} {
	return user.GetAll()
}


func (cntr UserController)Delete(id string) {
  user.Delete(id)
}


func (cntr UserController)Add(formData map[string]string) {
  if formData["Id"] != "" {
  	user.Update(formData)
  } else {
  	user.Create(formData)
  }
}



/*
func Edit() {
  fmt.Println("EDIT")
}
*/
func (cntr UserController)Get(id string)interface{} {
	return user.Get(id)
}
