package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shaviaditya/BasicGo/go_booksystem/pkg/models"
	"github.com/Shaviaditya/BasicGo/go_booksystem/pkg/utils"
	"github.com/go-delve/delve/pkg/dwarf/util"
	"github.com/gorilla/mux"
)

var newBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks:=models.GetBook()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request){
	vars :=mux.Vars(r)
	bookId:= vars["id"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("Error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name!="" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author!="" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication!="" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res,_ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)	
	w.Write(res)	
}    
