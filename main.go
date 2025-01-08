package main

import (
	"fmt"
	"net/http"
	"url-shortner/handlers"
)

func rootPageHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Url Shortening")
}

func main() {

	http.HandleFunc("/",rootPageHandler)
	http.HandleFunc("/shorten",handlers.ShortUrlHandler)
	http.HandleFunc("/redirect/", handlers.RedirectUrlHandler)
	http.HandleFunc("/metrics", handlers.MetricsHandler)

	fmt.Println("Server starting on port 8080")
	err:=http.ListenAndServe(":8080",nil)
	if(err!=nil){
		fmt.Println("Error on starting server")
	}
}
