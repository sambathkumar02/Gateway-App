package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func Get_Docker_version(c *fiber.Ctx) error {

	var connection *grpc.ClientConn
	var err error
	connection, err = grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		fmt.Print(err)
		return c.SendStatus(fiber.StatusNotFound)

	}
	defer connection.Close()
	grpc_client := NewDockerServiceClient(connection)

	client_data := Client{
		Client: c.IP() + ":" + c.Port(),
	}

	response, err := grpc_client.GetDockerVersion(context.Background(), &client_data)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	js, _ := json.Marshal(response)
	return c.Send(js)

}

func main() {
	app := fiber.New()

	app.Get("/version", Get_Docker_version)

	app.Listen(":8080")

}
