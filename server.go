package main
 
import (
            "database/sql"
            "fmt"
            _ "github.com/go-sql-driver/mysql"
            "github.com/labstack/echo"
            "github.com/labstack/echo/middleware"
            "net/http"
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
  Id     string `json:"id"`
  Name   string `json:"name"`
  Price string `json: "price"`
  Description    string `json : "description"`
  Quantity string `json: "quantity`
}
type Products struct {
  Products []Product `json:"product"`
}


db, err := sql.Open("mysql", "root:matti2527@tcp(127.0.0.1:3306)/barber")
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
 
	response := Product{Id: id, Name: name, Price: price, Description : description, Quantity: quantity}
	return c.JSON(http.StatusOK, response)
})

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}