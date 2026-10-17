package controllers

import "net/http"

func CreateUsers(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello, CreateUsers."))
}

func RetrieveUsers(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello, RetrieveUsers."))
}
