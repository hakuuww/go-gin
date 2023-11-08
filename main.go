package main

//import "fmt"
import (
	"jwtreactgo/database"
    "jwtreactgo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"	
)

func main() {
    database.Connect()


	app := fiber.New()

	//Allow origins default value is "*"
	//However, the application would run into CORS issues, even if i explicitly state 
	//AllowOrigins:"*"

	//However the application works when I use AllowOrigins: "http://localhost:3000"

	//This is likely due to the following


	//Browser Security: Browsers have security features that might affect CORS behavior. 
	//Some features are designed to prevent cross-origin requests for security reasons, 
	//and these features may be more restrictive when the * wildcard is used.
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost:3000",
	}))

    routes.Setup(app)

	app.Listen(":8000")
}
