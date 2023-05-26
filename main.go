package main

import (
	"code/competence/config"
	"code/competence/middleware"
	"code/competence/route"
	"time"
)

func main() {
	e := route.New()

	config.InitDB()

	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(":9000") }(lock)

	time.Sleep(1 * time.Millisecond)

	err := <-lock
	if err != nil {
		middleware.MakeLogEntry(nil).Panic("failed to start application")
	}
}
