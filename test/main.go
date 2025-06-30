package main

import (
	"log"

	"github.com/bakane6030/btk"
)

func main() {
	app, err := btk.Init("sigma", 640, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	<-app.Quit()
}
