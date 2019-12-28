package main

import (
  userController "./controller/userController"
	"./intfc/controllerIntfc"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var  tmp, error = template.ParseFiles("view/user/showAll.html", "view/current.html")

/**
  Метод куда приходят все запросы
*/
func main() {
  request := mux.NewRouter()
  request.HandleFunc("/{entity}", ParceHandler)
  request.HandleFunc("/{entity}/{action:create|save}", ParceHandler)
  request.HandleFunc("/{entity}/{action:view|edit|delete}/{id}", ParceHandler)

  http.Handle("/", request)

  fmt.Println("Start server")
  log.Fatal(http.ListenAndServe(":3030", nil))
}

/** Get needly controller  */
func getNeedlyController(entity string) controllerIntfc.ControllerIntfc {
	var currentController controllerIntfc.ControllerIntfc

  switch entity {
		 case "user" : currentController = userController.UserController{}
		 //case "post" : currentController = cntl.PostController{}    
  }

	return currentController
}

/**
  Parcing url params and call nessaccary function 
*/
func ParceHandler(responseWriter http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  entity := vars["entity"] 
  action := vars["action"] /* Type of action. What we want to do. For example Edit, Add, Delete, Show */
  id     := vars["id"]     /* Id of nessaccary object. We found object and do our "action" */

  contr := getNeedlyController(entity)
  
	if error != nil {
		fmt.Fprintf(responseWriter, error.Error())
	}
  /* 1. If we have "Id" it's mean we work this current entity  */
  if id != "" {
  	switch action{
  		case "delete":
  		      delete(contr, entity, id, responseWriter, request)
  			break
	case "edit":
      //delete(contr, entity, id, responseWriter, request)
  			break
  		case "view":
  		//	user := cntl.Get(id)
			//tmp.ExecuteTemplate(responseWriter, "current", user)
			break
  		default:
  			// Redirect to main page			
  	}

  	return
  }

  /* 2. New Entity */
  if action == "save" {
  	fmt.Println("SAVE")
	  request.ParseForm()
  	  fmt.Println(request.PostForm)
	  fmt.Println(request.Form)
	  //user := cntl.Add(request.Form)
	return
  }
  /* 3. Show all entitys */
  showAll(contr, entity,  responseWriter)
	
}


/* Отображаем все записи сущности */
func showAll(contr controllerIntfc.ControllerIntfc, entity string,  responseWriter http.ResponseWriter) {
  EntityArr := contr.ShowAll()
  fullTemplateName := entity + "showAll"
	tmp.ExecuteTemplate(responseWriter, fullTemplateName, EntityArr)
}


/* Удаление сущности */
func delete(contr controllerIntfc.ControllerIntfc, entity string, id string,  responseWriter http.ResponseWriter, request *http.Request) {
  contr.Delete(id)
  http.Redirect(responseWriter, request, "/"+entity, http.StatusSeeOther)
}


