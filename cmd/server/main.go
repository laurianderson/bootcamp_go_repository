package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/laurianderson/bootcamp_go_repository/cmd/server/handler"
	"github.com/laurianderson/bootcamp_go_repository/internal/product"
	"github.com/laurianderson/bootcamp_go_repository/pkg/store"
)

func main() {
	// Open database connection.
	databaseConfig := mysql.Config{
		User:   "root",
		Passwd: "",
		Addr:   "127.0.0.1:3306",
		DBName: "my_db",
	}

	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// Ping database connection.
	if err = database.Ping(); err != nil {
		panic(err)
	}

	println("Database connection established")

	//Change storage in json for storage in database
	storage := store.NewSqlStore(database)

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
