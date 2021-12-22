package main

import (
	"github.com/q-job/job/v1/internal/core"
)

func main() {
	app := core.NewCore()
	app.Run()
}
