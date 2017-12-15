package main

import (
	"agenda-api-first/service"
	"log"
	"github.com/emicklei/go-restful"
	"net/http"
)

func main() {
	wsContainer := restful.NewContainer()
    wsContainer.Router(restful.CurlyRouter{})
    service.MeetingRegister(wsContainer)
    service.UsersRegister(wsContainer)
    service.UserRegister(wsContainer)

    log.Printf("start listening on localhost:8080")
    server := &http.Server{Addr: ":8080", Handler: wsContainer}
    log.Fatal(server.ListenAndServe())
}