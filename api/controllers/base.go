package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialise (Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string){
	var err error
	if Dbdriver == "postgres"{
		DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB , err = gorm.Open(Dbdriver,DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		}else{
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
}