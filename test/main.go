package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bakane6030/btk"
)

func main() {
	app, err := btk.Init("sigma", 640, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-exit:
		slog.Info("SIGTERM received; exiting")
	case <-app.Quit():
		slog.Info("exit button clicked; exiting")
	}
}
