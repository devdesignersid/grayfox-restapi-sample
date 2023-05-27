package main

import (
	"fmt"
	"net/http"

	"github.com/devdesignersid/grayfox"
)

func main () {
  router := grayfox.NewRouter()
  
  // Adding a simple controller and registering the route
  controller := grayfox.NewController(*router)
  controller.Route("/", "GET", func(rw http.ResponseWriter, req *http.Request){
    fmt.Println("Hello World!")
  })

  // Adding goobye controller and registering "/goodbye" route
  goodByeController := grayfox.NewController(*router)
  goodByeController.Route("/goodbye", "GET", func(rw http.ResponseWriter, req *http.Request){
    fmt.Println("GoodBye World!")
  })

  // Registering Controllers
  router.Route("/", *controller)
  router.Route("/goodbye", *goodByeController)

  // Initiating the Grayfox App
  app := grayfox.New(*router)

  // Running the Grayfox app
  app.Run()
}
