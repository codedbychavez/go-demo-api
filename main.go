package main

import (
	"errors"
	"fmt"
	"go-demo-api/app/middleware"
	"go-demo-api/app/routing"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main()  {
	
	// Setup virtual environment
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	// Error checking for environment file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error, could not locate .env file", err)
		} else {
			fmt.Println("Error loading .env file", err)
		}
	}

	fiberConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}

	loggerConfig := fiberLogger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}

	app := fiber.New(fiberConfig)
	app.Use(fiberLogger.New(loggerConfig))
	app.Use(fiberRecover.New())


	baseRouter := app.Group("/api/v1")
	routing.CreatePublicRoutes(baseRouter)

	env, ok := viper.Get("ENVIRONMENT").(string)

	if !ok {
		err := errors.New("ENVIRONMENT not found")
		fmt.Println(err)
	}

	fmt.Println("Starting server via", "environment", env)

	startServer(app)
}

func startServer(app *fiber.App) {
	port, ok := viper.Get("PORT").(string)

	if !ok {
		err := errors.New("PORT not found")
		fmt.Println(err)
	}

	listenErr := app.Listen(":" + port)

	if listenErr != nil {
		fmt.Println("Error during listen", listenErr)
	}
}

