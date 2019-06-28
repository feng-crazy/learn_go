package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	value := req.FormValue("var1")
	fmt.Println("var1:", value)
	//fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
	fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", "go web", "ni ma you web")
	//http.HandlerFunc(HelloServer).ServeHTTP(w, req)
}

func main() {
	http.HandleFunc("/", HelloServer)
	//http.Handle("/", http.HandlerFunc(HelloServer))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}