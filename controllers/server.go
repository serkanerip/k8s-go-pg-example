package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/serkanerip/go-pg/models"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect
)

// Server for serving
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize setups database and router for server
func (server *Server) Initialize() {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("dbname"), os.Getenv("password"))

	server.DB, err = gorm.Open(os.Getenv("dbdriver"), DBURL)
	if err != nil {
		log.Fatalf("cannot connect to database %v", err)
	}
	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()
	server.initilizeRoutes()
}

// Run starts http server
func (server *Server) Run(addr string) {
	fmt.Printf("listening to %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
