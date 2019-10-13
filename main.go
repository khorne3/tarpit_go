package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	rtr := mux.NewRouter()
	http.Handle("/", rtr)

	// routes
	rtr.HandleFunc("/login", loginHandler)
	rtr.HandleFunc("/logout", logoutHandler)
	rtr.HandleFunc("/processOrder", processHandler)
	rtr.HandleFunc("/getOrderStatus", statusHandler)
	rtr.HandleFunc("/app", appHandler)
	rtr.HandleFunc("/init", initHandler)
	rtr.HandleFunc("/dbinit", dbinitHandler)
	//rtr.HandleFunc("/FileUploader", servicesHandler)
	// rtr.HandleFunc("/vulns", resourcesHandler)
	// rtr.HandleFunc("/insider", newResourceHandler)

	// listen to the web request
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
