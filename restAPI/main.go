package main

import (
	"fmt"
	"log"
	"net/http"
)


/**
Получаем на вход запрос, парсим его и возвращаем нужный ответ
 */
func parseUrl(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Println(request.Method)
	switch request.Method {
	case "GET":
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"message": "` + request.Method + `"}`))

	case "POST":
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"message": "` + request.Method + `"}`))

	default:
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte(`{"message": "method not found"}`))
	}

}

/**
Главная функция, которая всегда вызывается при инициализации
 */
func main() {
	http.HandleFunc("/", parseUrl)
	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
