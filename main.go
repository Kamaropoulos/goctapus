package main

import (
	"os"

	goctapus "github.com/Kamaropoulos/goctapus/core"
	"github.com/Kamaropoulos/goctapus/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args)

	goctapus.File("/", "public/index.html")
	goctapus.GET("/tasks", handlers.GetTasks(goctapus.Database))
	goctapus.PUT("/tasks", handlers.PutTask(goctapus.Database))
	goctapus.DELETE("/tasks/:id", handlers.DeleteTask(goctapus.Database))

	goctapus.Start()
}
