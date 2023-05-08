package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"

	"main.go/pkg/database"
	"main.go/pkg/models"
	"main.go/pkg/server"
)

func getDbConnectInfo() *models.DbConnectInfo {
	yfile, err := ioutil.ReadFile("./configs/database.yaml")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dbInfo := models.DbConnectInfo{}

	err = yaml.Unmarshal(yfile, &dbInfo)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &dbInfo
}

func main() {

	dbInfo := getDbConnectInfo()

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)
	fmt.Println(psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	dbConnect := database.New(db)
	router := gin.Default()

	s := server.New(dbConnect, router)
	s.Start()
}
