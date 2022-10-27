package main

import (
	"cd_platform/controller"
	"cd_platform/ext"
	"net/http"
	"sync"
)

func main() {
	ext.InitApp()

	e := controller.InitController()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go http.ListenAndServe(":8080", e)
	wg.Wait()
}
