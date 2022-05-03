package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloRoute(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found",http.StatusNotFound)
		return;
	}
	if r.Method != "GET" {
		http.Error(w, "Method Invalid",http.StatusNotFound)
	}
	fmt.Fprint(w, "Hello!");

}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err:= r.ParseForm(); err!=nil {
		fmt.Fprintf(w,"Parse error form : %v",err)
		return 
	}
	fmt.Fprintf(w,"Post request Successful\n");
	name := r.FormValue("Name")
	email := r.FormValue("Email")
	fmt.Fprintf(w,"Name : %s\n",name)
	fmt.Fprintf(w,"Email : %s\n",email)
}

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",helloRoute)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080",nil);
	fmt.Printf("%v\n",err)
	if err!=nil {
		log.Fatal(err)
	}


}