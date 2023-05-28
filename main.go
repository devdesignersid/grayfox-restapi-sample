package main

import (
	"fmt"
	"net/http"

	"github.com/devdesignersid/grayfox"
)

func main () {
  // Creating a Grayfox Router instance
  router := grayfox.NewRouter()

  // Adding a "/" (root) controller and it's handler.
  rootController := grayfox.NewController(*router)
  rootController.Route("/", "GET", func(rw http.ResponseWriter, req *http.Request){
    fmt.Println("Hello World!")
  })

  // Adding "goobye" controller and it' handler
  goodByeController := grayfox.NewController(*router)
  goodByeController.Route("/", "GET", func(rw http.ResponseWriter, req *http.Request){
    fmt.Println("GoodBye World!")
  })
  goodByeController.Route("/sing", "GET", func(rw http.ResponseWriter, req *http.Request){
    fmt.Println("I'm no good at goodbyes...")
  })


  // Registering Controllers
  router.Route("/", *rootController)
  router.Route("/goodbye", *goodByeController)

  // Initiating the Grayfox App
  app := grayfox.New(*router)

  // Running the Grayfox app
  app.Run()
}
