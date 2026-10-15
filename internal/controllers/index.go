package controllers

import (
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello, Index."))
}
