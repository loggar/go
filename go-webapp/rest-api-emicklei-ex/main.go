package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"

	userservice "github.com/loggar/go/go-webapp/rest-api-emicklei-ex/userservice"
)

func main() {
	restful.Add(userservice.New())
	log.Fatal(http.ListenAndServe(":18080", nil))
}
