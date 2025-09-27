package main

import (
	"dcdd_user_management_service/handlers"
	"dcdd_user_management_service/repositories"
	"dcdd_user_management_service/resolver"
	"dcdd_user_management_service/services"
	"dcdd_user_management_service/graph"
	"dcdd_user_management_service/helpers"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"
	

)

func main() {
	err := godotenv.Load()

	if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	db, err := helpers.GetGormDB()
    if err != nil {
        log.Fatal("Failed to connect to database: " + err.Error())
    }
    userRepository := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepository)
    resolver := resolvers.NewUserResolver(userService)

    mutationType := schema.NewMutationType(resolver)
	queryType := schema.NewQueryType(resolver)

	schema.InitSchema(queryType, mutationType)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Add any origins you need
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	e.POST("/upload-csv-file", handlers.UploadCSVFile)
	
	e.POST("/graphql", handlers.Handler)
	e.Logger.Fatal(e.Start(":8097"))
}

