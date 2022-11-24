package main

import (
	"cd_platform/controller"
	"cd_platform/ext"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ext.InitApp()

	e := controller.InitController()

	go func() {
		chann := make(chan os.Signal, 1000)
		signal.Notify(chann, os.Kill)

		for signall := range chann {
			fmt.Printf("%#v", signall)
		}
	}()

	http.ListenAndServe(":8080", e)
}
