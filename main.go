package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo EveryBody")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I,m Mohi Uddin")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method != "GET" {
		http.Error(w, "plz give me GET request", 400)
		return
	}
	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	preflightReq(w, r)
	if r.Method != "POST" {
		http.Error(w, "plz give me POST request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}
	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func preflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)
	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	pdr1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is Red, I love orange",
		Price:       100,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/4/43/Ambersweet_oranges.jpg",
	}
	pdr2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green, I love Green",
		Price:       350,
		ImgUrl:      "https://uttarakachabazar.com/wp-content/uploads/2022/12/Green_Apple-removebg-preview-copy.jpg",
	}
	pdr3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Eating two bananas daily can protect against these life-threatening diseases",
		Price:       30,
		ImgUrl:      "https://images.everydayhealth.com/images/diet-nutrition/bananas-nutrition-facts-1440x810.jpg?sfvrsn=5e5dc687_3",
	}

	productList = append(productList, pdr1)
	productList = append(productList, pdr2)
	productList = append(productList, pdr3)
}
