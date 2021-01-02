package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	type Product struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Price       string `json: "price"`
		Description string `json : "description"`
		Quantity    string `json: "quantity`
	}
	/*
		type Products struct {
			Products []Product `json:"product"`
		}*/

	db, err := sql.Open("mysql", "root:Guzzetta97^@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Routes
	e.GET("/", hello)

	e.GET("/products/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		fmt.Println(requested_id)
		var id string
		var name string
		var price string
		var description string
		var quantity string

		err = db.QueryRow("SELECT id,name, price, description, quantity_available FROM products WHERE id = ?", requested_id).Scan(&id, &name, &price, &description, &quantity)

		if err != nil {
			fmt.Println(err)
		}

		response := Product{Id: id, Name: name, Price: price, Description: description, Quantity: quantity}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/products", func(c echo.Context) error {
		fmt.Println("printntn")
		var products []Product
		//loading on results, scanning we need each key called
		results, err := db.Query("SELECT id,name, price, description, quantity_available FROM products")
		if err != nil {
			fmt.Println(err) //error handling - app
		}

		for results.Next() {
			var product Product
			// for each row, scan the result into our tag composite object
			err = results.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity)
			if err != nil {
				fmt.Println(err) //error handling - app
			}
			// and then print out attribute's tags
			log.Printf(product.Id, product.Name, product.Price, product.Description, product.Quantity)
			products = append(products, product)
		}

		response := products
		return c.JSON(http.StatusOK, response)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
