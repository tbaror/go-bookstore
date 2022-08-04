package routes

import (
	"github.com/gorilla/mux"
	"github.com/tbaror/go-bookstore/pkg/controllers"

)


var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreatBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}