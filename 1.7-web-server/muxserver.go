package main

import (
	"fmt"
	"http"
	"github.com/gorilla/mux"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)

}


func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
}