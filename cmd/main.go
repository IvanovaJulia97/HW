package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	localServer := server.MyServer(logger)

	if err := localServer.Start(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
