package main

import (
	"./controllers"
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	defaultMiddleware := middleware.DefaultMiddleware{}

	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())

	controllers.V1TaxesControllerHandler(router)

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)

	router.Run(serverString)

}
