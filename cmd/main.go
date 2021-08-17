package main

import (
	app "github.com/proggcreator/wb-openID"
	"github.com/proggcreator/wb-openID/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	handlers := handler.NewHandler()

	srv := new(app.Server)

	logrus.Print("Server starting...")
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
