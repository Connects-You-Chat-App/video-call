package main

import "connectsCall/internal/server"

func main() {
	// Run the server
	if err := server.Run(); err != nil {
		panic(err)
	}
}
