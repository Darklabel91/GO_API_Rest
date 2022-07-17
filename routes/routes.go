package routs

import (
	"github.com/Darklabel91/GO_API_Rest/controllers"
	"github.com/Darklabel91/GO_API_Rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	port              = ": 8000"
	main              = "/"
	allPersonalities  = "/api/personalidades"
	singlePersonality = "/api/personalidades/{ID}"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc(main, controllers.Home)
	router.HandleFunc(allPersonalities, controllers.ReturnAllPersonalities).Methods("Get")
	router.HandleFunc(allPersonalities, controllers.CreatePersonality).Methods("Post")
	router.HandleFunc(singlePersonality, controllers.ReturnPersonality).Methods("Get")
	router.HandleFunc(singlePersonality, controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc(singlePersonality, controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))

}
