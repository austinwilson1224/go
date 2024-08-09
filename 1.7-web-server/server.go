package main 

import (

"fmt"
"net/http"
"log"

"github.com/gorilla/mux"


)

func main() {
	fmt.Println("test")

	//http.HandleFunc("/test", handler) // each request calls handler 

	r := mux.NewRouter()

	r.HandleFunc("/", handler)


	log.Fatal(http.ListenAndServe("localhost:8082", nil))

}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
}



func handler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
}