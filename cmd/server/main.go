package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/apydox/apydox/pkg/common/bootstrap"
	"github.com/apydox/apydox/pkg/web"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed getting current working directory: %s", err.Error())
	}

	router.GET("/", web.IndexHandler)
	filesDir := fmt.Sprintf("%s/client/app/build/static", workingDir)
	router.ServeFiles("/static/*filepath", http.Dir(filesDir))
	// todo: serve the following files:
	// - asset-manifest.json
	// - favicon.ico
	// - manifest.json
	// - robots.txt

	port, err := getPortFromEnv()
	if err != nil {
		log.Fatal("invalid port value provided, the port must be an integer")
	}
	log.Printf("Listening on port %d ... \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func getPortFromEnv() (int, error) {
	port := os.Getenv("APYDOX_PORT")
	if port == "" {
		return bootstrap.ApyDoxDefaultPort, nil
	}
	return strconv.Atoi(port)
}
