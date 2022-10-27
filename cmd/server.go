package main

import (
	"cd_platform/controller"
	"cd_platform/ext"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ext.InitApp()

	e := controller.InitController()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go http.ListenAndServe(":8080", e)
	go handleSignal()
	wg.Wait()
}

func handleSignal() {
	sigchan := make(chan os.Signal)

	signal.Notify(sigchan)

	s := <-sigchan

	switch s {
	case syscall.SIGQUIT:
		clean()
		os.Exit(-1)
	case syscall.SIGKILL:
		clean()
		os.Exit(-1)
	case syscall.SIGTERM:
		clean()
		os.Exit(-1)
	}
}

func clean() {

}
