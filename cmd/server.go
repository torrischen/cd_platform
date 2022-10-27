package main

import (
	"cd_platform/controller"
	"cd_platform/ext"
)

func main() {
	ext.InitApp()

	e := controller.InitController()

	e.Run(":8080")
}
