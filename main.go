package main

import routes "github.com/guerraglucas/ginAPIwithGoroutine/routes"

func main() {
	router := routes.InitRoutes()

	router.Run(":8080")

}
