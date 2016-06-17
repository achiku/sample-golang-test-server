package main

import (
	"flag"
	"log"

	"github.com/achiku/sample-golang-test-server"
)

func main() {
	port := flag.String("p", "", "port number")
	flag.Parse()

	if *port != "" {
		s := client.NewMockServer(*port)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("service port number is not specified")
	}
}
