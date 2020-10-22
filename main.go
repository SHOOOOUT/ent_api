package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/shuto/ent-api/ent"
	"github.com/shuto/ent-api/handlers"
)

func main() {
	entOptions := []ent.Option{}

	entOptions = append(entOptions, ent.Debug())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mc := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "3306",
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	r.GET("/users", handlers.GetUsers)
	r.POST("/user/create", handlers.CreateUser)

	r.Run(":8080")
}
