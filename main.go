package main

import (
	"os"

	"github.com/MOZGIII/evans/adapter/controller"
	"github.com/MOZGIII/evans/meta"
)

func main() {
	os.Exit(controller.NewCLI(
		meta.AppName,
		meta.Version.String(),
		controller.NewBasicUI(),
	).Run(os.Args[1:]))
}
