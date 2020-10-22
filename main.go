package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yenchu/go-demo/handler"
	"github.com/yenchu/go-demo/repository"
)

func connectDb() {

	dbIP := os.Getenv("DB_IP")
	if dbIP == "" {
		log.Fatal("Database IP is required.")
	}

	portStr := os.Getenv("DB_PORT")
	var port int
	if portStr == "" {
		port = 3306
	} else {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal(err)
		}
		port = p
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		log.Fatal("Username is required.")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		log.Fatal("Password is required.")
	}

	config := repository.DbConfig{
		DbIP:     dbIP,
		DbPort:   port,
		DbName:   "demo",
		Username: username,
		Password: password,
	}

	repository.Connect(config)
}

func dataHandler(w http.ResponseWriter, req *http.Request) {

	handler := handler.NewDataHandler()

	if req.Method == "GET" {

		handler.FindAllData(w, req)

	} else if req.Method == "POST" {

		handler.CreateData(w, req)

	} else {

		http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)

	}
}

func main() {

	connectDb()

	http.HandleFunc("/data", dataHandler)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
