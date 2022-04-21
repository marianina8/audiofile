package main

import (
	metadataService "audiofile/services/metadata"
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 80, "Port for metadata service")
	flag.Parse()
	fmt.Printf("Starting API at http://localhost:%d\n", port)
	metadataService.Run(port)
}
