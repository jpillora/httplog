package main

import (
	"log"
	"os"

	"github.com/jpillora/httplog"
)

func main() {

	s, err := httplog.Dial("http", "1.1.1.1:80", httplog.LOG_INFO, "helloworld")
	if err != nil {
		println(err)
		os.Exit(1)
	}

	if err := s.Info("Hello world"); err != nil {
		log.Fatal(err)
	}
	log.Println("sent")
}
