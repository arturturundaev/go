package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
    "github.com/gorilla/mux"
    controller "./controller"
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
  request.HandleFunc("/{entity}", ParceHandler)
  request.HandleFunc("/{entity}/{action:create}", ParceHandler)
  request.HandleFunc("/{entity}/{action:view|edit|delete}/{id}", ParceHandler)

  http.Handle("/", request)

  fmt.Println("Start server")
  log.Fatal(http.ListenAndServe(":3030", nil))
}


/**
  Parcing url params and call nessaccary function 
*/
func ParceHandler(responseWriter http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  fmt.Println(vars)
  entity := vars["entity"] /* Type of object. For example User, Post, Article */
  action := vars["action"] /* Type of action. What we want to do. For example Edit, Add, Delete, Show */
  id     := vars["id"]     /* Id of nessaccary object. We found object and do our "action" */

  /* If we have "Id" it's mean we work this current entity  */
  if(id != nil) {

  }
fmt.Println(entity)  
fmt.Println(action)
fmt.Println(id)

}