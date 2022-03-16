package main

import (
	"authentication/internal/dependencies"
	"authentication/internal/repo"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := repo.NewMongoDB()
	if err != nil {
		log.Fatalln("Failed to connect to mongo")
	}

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app, err := dependencies.Inject(db)
	if err != nil {
		log.Fatalln("Failed to inject dependencies")
	}

	port := ":" + os.Getenv("PORT")

	if gin.Mode() == gin.ReleaseMode {
		log.Fatal(gateway.ListenAndServe(port, app))
	} else {
		log.Fatalln(http.ListenAndServe(port, app))
	}

}
