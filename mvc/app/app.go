package app

import (
	"net/http"

	"github.com/AkiraKane/microservices-golang/mvc/controllers"
)

func StartApp() {

	http.HandleFunc("/users", controllers.GetUser)

	http.ListenAndServe(":8080", nil)

}
