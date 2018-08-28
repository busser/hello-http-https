package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// HelloHandler writes a short message to w.
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}

// HelloConfiguration contains any necessary configuration.
type HelloConfiguration struct {
	httpPort          int
	httpsPort         int
	serverKey         string
	serverCertificate string
	uriPath           string
}

// Log logs the contents of conf.
func (conf *HelloConfiguration) Log() {
	log.Println("HTTP port:", conf.httpPort)
	log.Println("HTTPS port:", conf.httpsPort)
	log.Println("SSL key:", conf.serverKey)
	log.Println("SSL certificate:", conf.serverCertificate)
	log.Println("URI path:", conf.uriPath)
}

// ParseFlags parses CLI flags for any configuration parameters.
func ParseFlags() HelloConfiguration {
  var conf HelloConfiguration
	var helpWanted bool

	flag.BoolVar(&helpWanted, "help", false, "Prints this usage message")
  flag.IntVar(&conf.httpPort, "http-port", 80, "Port used to listen for HTTP requests")
  flag.IntVar(&conf.httpsPort, "https-port", 443, "Port used to listen for HTTPS requests")
  flag.StringVar(&conf.serverKey, "ssl-key", "server.key", "SSL key file for HTTPS server")
  flag.StringVar(&conf.serverCertificate, "ssl-certificate", "server.crt", "SSL certificate file for HTTPS server")
  flag.StringVar(&conf.uriPath, "uri-path", "/hello-world", "URI path to respond to")

	flag.Parse()

	if helpWanted {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0) // Exit after printing usage
	}

  return conf
}

// ListenHttpHttps starts HTTP and HTTPS servers based on conf.
func ListenHttpHttps(conf HelloConfiguration) {
  // wg is used to stop the function from returning right after starting the
  // HTTP and HTTPS servers.
	var wg sync.WaitGroup
	wg.Add(2)

	http.HandleFunc(conf.uriPath, HelloHandler)

  // Start the HTTP server.
	go func() {
		defer wg.Done()
    log.Println("Listening for HTTP requests...")
		err := http.ListenAndServe(fmt.Sprint(":", conf.httpPort), nil)
		if err != nil {
			log.Fatal(fmt.Sprint("Could not listen on port ", conf.httpPort, ": "), err)
		}
	}()

  // Start the HTTPS server.
	go func() {
		defer wg.Done()
    log.Println("Listening for HTTPS requests...")
		err := http.ListenAndServeTLS(fmt.Sprint(":", conf.httpsPort), conf.serverCertificate, conf.serverKey, nil)
		if err != nil {
			log.Fatal(fmt.Sprint("Could not listen on port ", conf.httpsPort, ": "), err)
		}
	}()

	wg.Wait()
}

func main() {
  conf := ParseFlags()
	conf.Log()
	ListenHttpHttps(conf)
}
