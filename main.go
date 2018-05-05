package main

import (
	"os"

	"github.com/Kamaropoulos/goctapus/core"
	"github.com/Kamaropoulos/goctapus/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args, "debug")

	goctapus.ConnectDB("goapp")
	goctapus.Migrate(goctapus.Databases["goapp"], "./models/tasks.sql")

	goctapus.File("/", "public/index.html")
	goctapus.GET("/tasks", handlers.GetTasks(goctapus.Databases["goapp"]))
	goctapus.PUT("/tasks", handlers.PutTask(goctapus.Databases["goapp"]))
	goctapus.DELETE("/tasks/:id", handlers.DeleteTask(goctapus.Databases["goapp"]))

	goctapus.Start()
}
