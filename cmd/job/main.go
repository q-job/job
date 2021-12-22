package main

import (
	"github.com/q-job/job/internal/core/v1"
)

func main() {
	app := core.NewCore()
	app.Run()
}
