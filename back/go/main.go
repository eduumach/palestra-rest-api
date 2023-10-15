package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Product struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Preco     float32 `json:"preco"`
	Vendido   bool   `json:"vendido"`
}

type ProductCreateAndUpdate struct {
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Preco     float32 `json:"preco"`
	Vendido   bool   `json:"vendido"`
}

type ProductPartialUpdate struct {
	Vendido   bool   `json:"vendido"`
}

func main() {
	//connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if it doesn't exist

	if err != nil {
		log.Fatal(err)
	}

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/produtos", getProducts(db)).Methods("GET")
	router.HandleFunc("/produtos/{id}", getProduct(db)).Methods("GET")
	router.HandleFunc("/produtos", createProduct(db)).Methods("POST")
	router.HandleFunc("/produtos/{id}", updateProduct(db)).Methods("PUT")
	router.HandleFunc("/produtos/{id}", partialUpdateProduct(db)).Methods("PATCH")
	router.HandleFunc("/produtos/{id}", deleteProduct(db)).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":8050", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Backend-Type", "Golang")
		next.ServeHTTP(w, r)
	})
}

// get all products
func getProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM produtos")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		products := []Product{}
		for rows.Next() {
			var u Product
			if err := rows.Scan(&u.ID, &u.Nome, &u.Descricao, &u.Preco, &u.Vendido); err != nil {
				log.Fatal(err)
			}
			products = append(products, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(products)
	}
}

// get product by id
func getProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var u Product
		row := db.QueryRow("SELECT * FROM produtos WHERE id = $1", params["id"])
		if err := row.Scan(&u.ID, &u.Nome, &u.Descricao, &u.Preco, &u.Vendido); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// create product
func createProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u Product
		var u1 ProductCreateAndUpdate
	
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("INSERT INTO produtos (nome, descricao, preco) VALUES ($1, $2, $3)", u.Nome, u.Descricao, u.Preco)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u1)
	}
}

// update product
func updateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var u ProductCreateAndUpdate
	
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("UPDATE produtos SET nome = $1, descricao = $2, preco = $3 WHERE id = $4", u.Nome, u.Descricao, u.Preco, params["id"])
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// partial update product
func partialUpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var u ProductPartialUpdate
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("UPDATE produtos SET vendido = $1 WHERE id = $2", u.Vendido, params["id"])
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// delete product
func deleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		_, err := db.Exec("DELETE FROM produtos WHERE id = $1", params["id"])
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
