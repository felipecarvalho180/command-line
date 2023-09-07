package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

const startLine = "Start line"

func main() {
	fmt.Println(startLine)

	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
