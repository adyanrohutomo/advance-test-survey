package main

import (
	"advance-test-survey/db"
	"advance-test-survey/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
