package service

import (
	"log"
	"net/http"
)

// StartWebServer =)
func StartWebServer(port string) {
	log.Println("Starting http service at: " + port)

	r := NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("an error occured starting http server at port" + port)
		log.Println("error: " + err.Error())
	}
}
