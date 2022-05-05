package routes
import (
	"github.com/gorilla/mux"
	"github.com/Shaviaditya/BasicGo/go_booksystem/pkg/controllers"
)

var BookRoutes = func (router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}",controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}",controllers.UpdateBook).Methods("PUT")
}