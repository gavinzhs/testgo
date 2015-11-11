package main

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestBaseroute(t *testing.T) {
	//	initDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	host := os.Getenv("HOST")
	log.Println("host:", host)
	if host == "" {
		host = "localhost"
		log.Println("是空得 我需要重新设置host值")
	}

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/first/", sayhelloName)
	http.HandleFunc("/second/", sayhelloName)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
