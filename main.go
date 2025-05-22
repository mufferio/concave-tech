package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mufferio/concave-tech/controllers"
	"github.com/mufferio/concave-tech/repositories"
	"github.com/mufferio/concave-tech/services"
	"github.com/mufferio/concave-tech/utilities"
)

func main() {
	db := utilities.NewDB()
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	svc := services.NewUserService(repo)
	ctrl := controllers.NewUserController(svc)

	r := gin.Default()
	ctrl.Register(r)

	log.Println("⇢  http://localhost:8080/users")
	log.Fatal(r.Run(":8080"))
}
