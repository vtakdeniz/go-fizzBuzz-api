package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type response struct {
	Arr []string `json:"fizzBuzz"`
}

func setRoutes(app *fiber.App) {
	app.Get("fizzbuzz/:count", fizzBuzz)
}

func fizzBuzz(ctx *fiber.Ctx) error {
	num, err := strconv.Atoi(ctx.Params("count"))
	if err != nil {
		panic("Error converting url parameter to integer")
	}

	var response response

	for i := 1; i < num; i++ {
		if i%15 == 0 {
			response.Arr = append(response.Arr, "fizz buzz")
		} else if i%5 == 0 {
			response.Arr = append(response.Arr, "buzz")
		} else if i%3 == 0 {
			response.Arr = append(response.Arr, "fizz")
		} else {
			response.Arr = append(response.Arr, strconv.Itoa(i))
		}
	}
	return ctx.JSON(response)
}

func main() {
	port := ":9000"
	app := fiber.New()
	app.Use(cors.New())
	setRoutes(app)
	err := app.Listen(port)
	if err != nil {
		log.Fatal("Port is in use")
	}
}
