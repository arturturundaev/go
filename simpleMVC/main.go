package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)

type Message struct {
	Text string
}

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
			fmt.Println("eror")
			fmt.Fprintf(responseWriter, error.Error())
		} else  {
			fmt.Println("success")
			
		
			msg := Message{Text:"Hello, world!"}
			MessageArr = append(MessageArr, msg)
			fmt.Println(MessageArr)
			tmp.ExecuteTemplate(responseWriter, "index", MessageArr)
		}
	}
} 

func main() {
	http.HandleFunc("/", parseUrl)
	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":3030", nil))
}