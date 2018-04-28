package main

import (
	"os"

	goctapus "github.com/Kamaropoulos/goctapus/core"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args)
	goctapus.AddEndpoints()
	goctapus.Start()
}
