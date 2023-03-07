package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rizkyfazri23/dripay/routes"
)

func main() {
	dbHost := "localhost"
	dbPort := "54321"
	dbName := "final_project"
	dbUser := "postgres"
	dbPassword := "dendi16"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	g := gin.Default()

	routes.GatewayRoutes(g, db)

	g.Run(":8080")
}
