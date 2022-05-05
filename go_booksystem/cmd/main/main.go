package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Shaviaditya/BasicGo/go_booksystem/pkg/routes"
)


func main(){
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9050",r)) 
}