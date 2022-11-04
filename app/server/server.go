package server

import (
	"fmt"
	"net/http"
	"strconv"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")
	goly, err := models.FindByGolyUrl(golyUrl)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": " could not find goly in Db" + err.Error(),
			},
		)
	}
	// get stats
	goly.Clicked += 1
	err = models.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("Error Updating %v\n", err.Error())
	}
	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllGolies(c *fiber.Ctx) error {
	golies, err := models.GetAllGolies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not parse id" + err.Error(),
			},
		)
	}

	goly, err := models.GetGoly(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "error could not retrive goly from the database",
			},
		)
	}

	return c.Status(http.StatusOK).JSON(goly)

}

func createGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly models.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": " could not parse json" + err.Error(),
			},
		)
	}

	if goly.Random {
		goly.Goly = utils.Random(8)
	}
	err = models.CreateGoly(goly)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not create Goly in Database",
			},
		)
	}

	return c.Status(http.StatusOK).JSON(&goly)
}

func updateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly models.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": " could not parse json" + err.Error(),
			},
		)
	}
	err = models.UpdateGoly(goly)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": " could not parse json" + err.Error(),
			},
		)
	}

	return c.Status(http.StatusOK).JSON(&goly)
}

func deleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not find id" + err.Error(),
			},
		)
	}
	err = models.DeleteGoly(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not delete goly in db" + err.Error(),
			},
		)
	}
	return c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Successfully delted goly",
		},
	)
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("at golly")
	})

	router.Get("/r/:redirect", redirect)
	router.Get("/goly", getAllGolies)
	router.Get("/goly/:id", getGoly)
	router.Post("/goly", createGoly)
	router.Put("/goly", updateGoly)
	router.Delete("/goly/:id", deleteGoly)

	router.Listen(":8080")
}
