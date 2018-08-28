package main

import (
	"log"
)

func main() {
	// Enrich logs: add microseconds to times.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Get configuration from CLI flags.
	conf := ParseFlags()
	conf.Log()

	// Listen for HTTP and HTTPS requests
	ListenHttpHttps(conf)
}
