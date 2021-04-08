package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-project/pkg/transport"
	"net/http"
	"os"
)

const serverUrl = ":8000"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the serve")
	r := transport.Router()
	fmt.Println(http.ListenAndServe(serverUrl, r))
}
