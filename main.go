package main

import (
	"backend-task-eventori/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
