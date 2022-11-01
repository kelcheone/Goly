package main

import (
	"url-shortener/models"
	"url-shortener/server"
)

func main() {
	models.Setup()

	server.SetupAndListen()
}
