package main

import (
	"ApolloUtils/internal/app"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load()

	utilsApp := app.New()

	contacts, err := utilsApp.FindSponsors("Unstop")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range contacts {
		fmt.Println("---")
		fmt.Println(c.Name)
		fmt.Println(c.Title)
		fmt.Println(c.ApolloURL)
	}
}
