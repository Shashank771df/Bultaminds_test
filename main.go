package main

import (
	db "transactions/app/database"
	router "transactions/router"
)

func main() {
	db.Init()
	e := router.Init()
	e.Run(":1323")
}
