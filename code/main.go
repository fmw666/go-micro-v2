package main

import "app/router"

func main() {
	r := router.Router()
	r.Run(":8080")
}
