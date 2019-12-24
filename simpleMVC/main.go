package main

import (
	cntl "./controller"
	intfc "./intfc"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)


/* Объект Сообщение. Содержит только одно поле - Text */
type Message struct {
	Text string
	Id int
}

var increment int

/* Массив, каждый элемент который должен быть Message */
var MessageArr = make([]Message, 0)

/**
@todo future определить какой тип запроса - 
если json, то и возращать json
*/
func parseUrl(responseWriter http.ResponseWriter, request *http.Request) {
	
	switch request.Method {
	case "GET":
		tmp, error := template.ParseFiles("view/index.html")

		if error != nil {
			fmt.Fprintf(responseWriter, error.Error())
		} else  {		
			//currentId := path.Base(request.URL.Path)
			//if currentId != nil {
				//tmp.ExecuteTemplate(responseWriter, "current", MessageArr[currentId])
			//}
			msg := Message{Text:"Hello, world!", Id:increment}
			increment++
			MessageArr[increment] = msg
			tmp.ExecuteTemplate(responseWriter, "index", MessageArr)
		}
	}
} 


/**
  Метод куда приходят все запросы
*/
func main() {
  request := mux.NewRouter()
  request.HandleFunc("/user", ParceHandler)
  request.HandleFunc("/user/{action:create|save}", ParceHandler)
  request.HandleFunc("/user/{action:view|edit|delete}/{id}", ParceHandler)

  http.Handle("/", request)

  fmt.Println("Start server")
  log.Fatal(http.ListenAndServe(":3030", nil))
}

/** Get needly controller  */
func GetNeedlyController(entity string) intfc.Intfc {
	var currentController intfc.Intfc

	if entity == "user" {
		currentController = cntl.UserController{}
	}

	return currentController
}

/**
  Parcing url params and call nessaccary function 
*/
func ParceHandler(responseWriter http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  action := vars["action"] /* Type of action. What we want to do. For example Edit, Add, Delete, Show */
  id     := vars["id"]     /* Id of nessaccary object. We found object and do our "action" */

  tmp, error := template.ParseFiles("view/showAll.html", "view/current.html")
	if error != nil {
		fmt.Fprintf(responseWriter, error.Error())
	}
  /* 1. If we have "Id" it's mean we work this current entity  */
  if id != "" {
  	switch action{
  		case "delete":
  			cntl.Delete(id)
  			break
	case "edit":
  			cntl.Edit()
  			break
  		case "view":
  			user := cntl.Get(id)
			tmp.ExecuteTemplate(responseWriter, "current", user)
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
  contr := GetNeedlyController("user")
	EntityArr := contr.ShowAll()
	tmp.ExecuteTemplate(responseWriter, "showAll", EntityArr)



}


