package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huemmerj/TourneyShare/db"
	"github.com/huemmerj/TourneyShare/handlers"
	"os"
)

func SetHeader(header, value string, handle http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set(header, value)
		handle.ServeHTTP(w, req)
	}
}
func main() {
	r := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitMongoDB()
	fs := http.FileServer(http.Dir("./dist"))
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))
	// start server and log error

	r.Handle("/", handlers.HomeHandler())
	r.Handle("/addTournament", handlers.AddTournamentHandler()).Methods("GET")
	r.HandleFunc("/addTournament", handlers.AddTournamentSubmitHandler).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println(port)
	port = fmt.Sprintf(":%s", port)
	fmt.Printf("Server drunning on Port %s", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println(err)
	}
}
