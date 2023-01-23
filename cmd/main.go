package main

import (
	"github.com/VadimGossip/conversionWriter/internal/app"
	"time"
)

func main() {
	convWriter := app.NewApp("Conversion writer", time.Now())
	convWriter.Run()
}
