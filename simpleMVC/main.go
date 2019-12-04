package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)


/* Объект Сообщение. Содержит только одно поле - Text */
type Message struct {
	Text string
}


/* Массив, каждый элемент который должен быть Message */
var MessageArr []Message

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
			msg := Message{Text:"Hello, world!"}
			MessageArr = append(MessageArr, msg)
			tmp.ExecuteTemplate(responseWriter, "index", MessageArr)
		}
	}
} 


/**
  Метод куда приходят все запросы
*/
func main() {
	http.HandleFunc("/", parseUrl)
	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":3030", nil))
}