package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mamaz/romanserver/roman-numerals"
)

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	urlElements := strings.Split(request.URL.Path, "/")
	number, err := strconv.Atoi(urlElements[2])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 internal server error"))
		return
	}

	if number == 0 || number > 10 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
	} else {
		fmt.Fprintf(writer, "%q", html.EscapeString(romanNumerals.Numerals[number]))
	}
}

func main() {
	http.HandleFunc("/", httpHandler)
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
