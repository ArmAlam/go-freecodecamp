package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main()  {
	fmt.Println("Hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if(portString == "") {
		log.Fatal("PORT is not set in env")
	}

	fmt.Println("PORT is ", portString)
}