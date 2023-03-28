package main

import (
	"hamster-paas/pkg/api"
	"hamster-paas/pkg/initialization"
	"os"
)

func main() {
	initialization.Init()
	api.Serve(os.Getenv("PORT"))
}
