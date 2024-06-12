package main

import (
	"chatter/controller"
	"chatter/middleware"
	"chatter/router"
	"chatter/server"
)

func setup() server.Info {
	return server.Info{
		Hostname:        "",
		UseHTTP:         true,
		UseHTTPS:        false,
		RedirectToHTTPS: false,
		HTTPPort:        80,
		HTTPSPort:       443,
	}
}

func main() {
	handler := middleware.SetUp(router.Instance())

	controller.Load()
	// Start the HTTP and HTTPS listeners
	server.Run(
		handler, // HTTP handler
		handler, // HTTPS handler
		setup(), // Server settings
	)
}
