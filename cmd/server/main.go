package main

import (
	"fmt"
	"github.com/1Str1de1/Avito-backend-task/internal/app/server"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Info: .env file not found, using environment variables")
	}
	conf := server.NewConfig()

	s := server.New(conf)

	go s.MustStart(conf)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	stopSignal := <-stop

	fmt.Println(stopSignal)

	s.Stop()
}
