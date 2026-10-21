package main

import (
	"qrix-runtime/internal/app"
)

func main() {
	a := app.NewApp()

	if err := a.Start(); err != nil {
		//todo: log error and exit
	}
}
