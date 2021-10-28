package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api", api)
	app.Get("/api2", api2)

	app.Listen(":8001")
}

func init() {

	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 4,
		ErrorPercentThreshold:  50,
		SleepWindow:            15000,
	})

	hystrix.ConfigureCommand("api2", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 4,
		ErrorPercentThreshold:  50,
		SleepWindow:            15000,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8002", hystrixStreamHandler)

}

func api(c *fiber.Ctx) error {

	output := make(chan string, 1)
	hystrix.Go("api", func() error {

		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		output <- msg

		return nil
	}, func(err error) error {
		fmt.Println(err)

		return err
	})

	out := <-output

	return c.SendString(out)
}

func api2(c *fiber.Ctx) error {

	output := make(chan string, 1)
	hystrix.Go("api2", func() error {

		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		output <- msg

		return nil
	}, func(err error) error {
		fmt.Println(err)

		return nil
	})

	out := <-output

	return c.SendString(out)
}
