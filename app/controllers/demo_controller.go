package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Define the controller
type DemoController struct {
	response string
}

func NewDemoController(response string) DemoController {
	return DemoController{
		response: response,
	}
}

// FindRandom Demo returns some random string 
func (ctrl DemoController) FindRandom(c *fiber.Ctx) error {
	return c.SendString("I am returned from FindRandom method")
}



