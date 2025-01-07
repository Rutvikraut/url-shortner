package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type URL struct {
	ID           string `json:"id"`
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}

var urlDB = make(map[string]URL)

func generateShortUrl(originalUrl string) string {
	hasher:= md5.New()
	hasher.Write([]byte(originalUrl))
	data:=hasher.Sum(nil)
	
	hash:= hex.EncodeToString(data)

	fmt.Println(hash[:6])
	return hash[:6]
}

func createShortUrl(originalURL string) string {
	shortUrl:=generateShortUrl(originalURL)
	id:=shortUrl
	urlDB[id] = URL{
		ID: id,
		OriginalURL: originalURL,
		ShortenedURL: shortUrl,
	}

	return shortUrl
}

func getUrl(id string) (URL,error){
	url,ok :=urlDB[id]
	if(!ok){
		return URL{}, errors.New("URL Not Found")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get method")
	fmt.Fprintf(w,"Hello World")
}

func shortUrlHandler(w http.ResponseWriter, r *http.Request){
	var data struct{
		URL string `json:"url"`
	}

	err:=json.NewDecoder(r.Body).Decode(&data)

	if(err!=nil){
		http.Error(w,"Invalid Request",http.StatusBadRequest)

		return
	}

	shortUrl_:=createShortUrl(data.URL)

	response:= struct{
		ShortURL string `json:"short_url`
	}{ShortURL: shortUrl_}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectUrlHandler(w http.ResponseWriter, r *http.Request){
	id:=r.URL.Path[len("/redirect/"):]
	url, err := getUrl(id)

	if(err!=nil){
		http.Error(w,"Invalid Request", http.StatusNotFound)
	}

	http.Redirect(w,r,url.OriginalURL,http.StatusFound)
}

func main() {

	http.HandleFunc("/",handler)
	http.HandleFunc("/shorten",shortUrlHandler)
	http.HandleFunc("/redirect/", redirectUrlHandler)

	fmt.Println("Server starting on port 8080")
	err:=http.ListenAndServe(":8080",nil)
	if(err!=nil){
		fmt.Println("Error on starting server")
	}
}
