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

type Employee struct {
  Id     string `json:"id"`
  Name   string `json:"employee_name"`
  Salary string `json: "employee_salary"`
  Age    string `json : "employee_age"`
}
type Employees struct {
  Employees []Employee `json:"employee"`
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

e.GET("/employee/:id", func(c echo.Context) error {
	requested_id := c.Param("id")
	fmt.Println(requested_id)
	var name string
	var id string
	var salary string
	var age string
 
	err = db.QueryRow("SELECT id,employee_name, employee_age, employee_salary FROM employee WHERE id = ?", requested_id).Scan(&id, &name, &age, &salary)
 
	if err != nil {
				fmt.Println(err)
	}
 
	response := Employee{Id: id, Name: name, Salary: salary, Age: age}
	return c.JSON(http.StatusOK, response)
})

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}