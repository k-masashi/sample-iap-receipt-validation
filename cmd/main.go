package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"sample.com/sample"
)

func main() {
	funcframework.RegisterHTTPFunction("/", sample.SampleGet)
	funcframework.RegisterHTTPFunction("/verify", sample.VerifyReceipt)
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
