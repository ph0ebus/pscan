package main

import (
	"pscan/plugins"
	"pscan/web/service"
)

func main() {
	logger := plugins.LogInit()
	service.RunService(logger)
}
