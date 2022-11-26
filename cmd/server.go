package main

import (
	"cd_platform/controller"
	"cd_platform/ext"
	"net/http"
)

func main() {
	ext.InitApp()

	e := controller.InitController()

	http.ListenAndServe(":8080", e)
}
