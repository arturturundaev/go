package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)

func parseUrl(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	switch request.Method {
	case "GET":
		tmp, error := template.ParseFiles("view/index.html")

		if error != nil {
			fmt.Println("eror")
			fmt.Fprintf(responseWriter, error.Error())
		} else  {
			fmt.Println("success")
			tmp.ExecuteTemplate(responseWriter, "index", nil)
		}
	}
} 

func main() {
	http.HandleFunc("/", parseUrl)
	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":3030", nil))
}