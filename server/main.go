package main

import (
	"log"

	"github.com/Sharykhin/go-users-grpc/server/handler"
)

func main() {
	log.Fatal(handler.ListenAndServe())
}
