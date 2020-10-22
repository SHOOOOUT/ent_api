package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/shuto/ent-api/ent"
	"github.com/shuto/ent-api/ent/user"
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

	r := echo.New()
	r.GET("/users", func(c echo.Context) error {
		eq := client.User.Query()
		entries := eq.AllX(context.Background())
		return c.JSON(http.StatusOK, entries)
	})

	r.GET("/user/:id", func(c echo.Context) error {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
		eq := client.User.
			Query().
			Where(user.IDEQ(int(id)))
		u := eq.OnlyX(context.Background())
		return c.JSON(http.StatusOK, u)
	})

	r.Start(":8080")
}
