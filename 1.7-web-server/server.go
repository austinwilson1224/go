package main 

import (

"fmt"
"net/http"
"log"

)

func main() {
	fmt.Println("test")

	http.HandleFunc("/", handler) // each request calls handler 


	log.Fatal(http.ListenAndServe("localhost:8082", nil))

}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
}