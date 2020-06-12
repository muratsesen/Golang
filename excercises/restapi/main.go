// //Tutorial link: https://www.youtube.com/watch?v=SonwZ6MF5BE
// //Book application

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	// "math/rand"
// 	// "strconv"
// 	"github.com/gorilla/mux"
// )
// //Book Struct
// type Book struct{
// 	ID 		string `json:"id"`
// 	Isbn 	string `json:"isbn"`
// 	Title 	string `json:"title"`
// 	Author *Author `json:"author"`
// }
// //Author Struck
// type Author struct{
// 	FirstName 	string `json:"firstname"`
// 	LastName 	string `json:"lastname"`
// }

// //Get Books
// func getBooks(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(books)
// }
// //Get Single Book
// func getBook(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	params := mux.Vars(r) //get params
// 	//Loop through books to find the requested
// 	for _,item:=range books{
// 		if item.ID == params["id"]{
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }
// //Create a New Book
// func createBook(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = "4"
// 	books = append(books,book)
// 	json.NewEncoder(w).Encode(book)
// }
// //
// func updateBook(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	params := mux.Vars(r) 
// 	for index, item :=range books{
// 		if item.ID == params["id"]{
// 			//first delete existing one
// 			books = append(books[:index],books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books,book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// }
// //
// func deleteBook(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	params := mux.Vars(r)
// 	for index,item:=range books{
// 		if item.ID == params["id"]{
// 			books = append(books[:index],books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// var books []Book

// func main(){
// 	//Init router
// 	r := mux.NewRouter()

// 	//Mock data - @todo - implement DB
// 	books = append(books, Book{ID:"1",Isbn:"44567",Title:"Booke One",Author:&Author{FirstName:"John", LastName:"Doe"}})
// 	books = append(books, Book{ID:"2",Isbn:"88561",Title:"Booke Two",Author:&Author{FirstName:"John", LastName:"Doe"}})
// 	books = append(books, Book{ID:"3",Isbn:"33565",Title:"Booke Three",Author:&Author{FirstName:"Magni", LastName:"Ficent"}})

// 	//Route Handlers / Endpoints
// 	r.HandleFunc("/api/books",getBooks).Methods("GET")
// 	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
// 	r.HandleFunc("/api/books",createBook).Methods("POST")
// 	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
// 	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":8000",r))
// }

//////////////////////////////////////////////////////////////////////////////////
//This is my turn to practice go api
//////////////////////////////////////////////////////////////////////////////////

package main

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

///////////////////////////////////////////////////////
//-- Structs
///////////////////////////////////////////////////////
type Car struct{
	ID string `json:"id"`
	Shasis string `json:"shasis"`
	Make *Make `json:"make"` 
	Model *Model `json:"model"`
}

type Make struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
}
type Model struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

///////////////////////////////////////////////////////
//-- Methods
///////////////////////////////////////////////////////
func getCars(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(cars)
}
func getCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var params = mux.Vars(r)
	for _,item :=range cars{
		if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}
func saveCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//get new car 
	var car Car
	_=json.NewDecoder(r.Body).Decode(&car)
	car.ID = "5"

	//add list
	cars = append(cars,car)
	json.NewEncoder(w).Encode(cars)
}
func updateCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item :=range cars{

		if item.ID == params["id"]{

			//first delete existing one
			cars = append(cars[:index],cars[index+1:]...)
			var car Car
			_=json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars,car)
			json.NewEncoder(w).Encode(cars)
		return
		}
	}
}
func deleteCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item :=range cars{

		if item.ID == params["id"]{

			// delete existing one
			cars = append(cars[:index],cars[index+1:]...)
			json.NewEncoder(w).Encode(cars)
		return
		}
	}
}

///////////////////////////////////////////////////////
//-- Globals
///////////////////////////////////////////////////////
var cars []Car
///////////////////////////////////////////////////////
//-- Main
///////////////////////////////////////////////////////
func main(){
	r := mux.NewRouter()

	//MOCK Data
	cars = append(cars,Car{ID:"1",Shasis:"12344",Make: &Make{ID:"1",Name:"Honda",Country:"Japan"},Model: &Model{ID:"1",Name:"Jazz"}})
	cars = append(cars,Car{ID:"2",Shasis:"13333",Make: &Make{ID:"1",Name:"Honda",Country:"Japan"},Model: &Model{ID:"2",Name:"Civic"}})
	cars = append(cars,Car{ID:"3",Shasis:"14444",Make: &Make{ID:"1",Name:"Honda",Country:"Japan"},Model: &Model{ID:"3",Name:"City"}})

	r.HandleFunc("/api/cars",getCars).Methods("GET")
	r.HandleFunc("/api/cars/{id}",getCar).Methods("GET")
	r.HandleFunc("/api/cars",saveCar).Methods("POST")
	r.HandleFunc("/api/cars/{id}",updateCar).Methods("PUT")
	r.HandleFunc("/api/cars/{id}",deleteCar).Methods("DELETE")

	http.ListenAndServe(":8000",r)
}