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

type User struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Preco     string `json:"preco"`
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
	router.HandleFunc("/produtos", getUsers(db)).Methods("GET")
	router.HandleFunc("/produtos/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/produtos", createUser(db)).Methods("POST")
	router.HandleFunc("/produtos/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/produtos/{id}", partialUpdateUser(db)).Methods("PATCH")
	router.HandleFunc("/produtos/{id}", deleteUser(db)).Methods("DELETE")

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

// get all users
func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM produtos")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Nome, &u.Descricao, &u.Preco, &u.Vendido); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(users)
	}
}

// get user by id
func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var u User
		row := db.QueryRow("SELECT * FROM produtos WHERE id = $1", params["id"])
		if err := row.Scan(&u.ID, &u.Nome, &u.Descricao, &u.Preco, &u.Vendido); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// create user
func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("INSERT INTO produtos (nome, descricao, preco) VALUES ($1, $2, $3)", u.Nome, u.Descricao, u.Preco)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// update user
func updateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var u User
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("UPDATE produtos SET nome = $1, descricao = $2, preco = $3 WHERE id = $4", u.Nome, u.Descricao, u.Preco, params["id"])
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// partial update user
func partialUpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var u User
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("UPDATE produtos SET vendido = $1 WHERE id = $2", u.Vendido, params["id"])
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// delete user
func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		_, err := db.Exec("DELETE FROM produtos WHERE id = $1", params["id"])
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
