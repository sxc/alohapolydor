package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/sxc/alohapolydor/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	appv1 := app.Group("api/v1")

	appv1.Get("/user", api.HandleGetUsers)
	appv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
