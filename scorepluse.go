package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Neeraj1997-dev/scorepluse-pipeline-services/routes"
	"github.com/joho/godotenv"
)

var err error

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	bootstrap()
}

func bootstrap() {
	defer restart()
	err = routes.App.Listen(fmt.Sprintf(":%s", os.Getenv("SERVERPORT")))
	if err != nil {
		panic(err)
	}
}

func restart() {
	if err != nil {
		fmt.Println("waiting for 5 second to restart app", err)
		time.Sleep(time.Second * 5)
		bootstrap()
	}
}
