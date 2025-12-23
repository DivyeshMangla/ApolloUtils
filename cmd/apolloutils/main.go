// ApolloUtils provides a CLI for finding sales and marketing contacts via Apollo.io.
package main

import (
	"ApolloUtils/internal/app"
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	_ = godotenv.Load()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter company name: ")
	companyName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	companyName = strings.TrimSpace(companyName)

	if companyName == "" {
		log.Fatal("Company name cannot be empty")
	}

	utilsApp := app.New()

	contacts, err := utilsApp.FindSponsors(companyName)
	if err != nil {
		log.Fatal(err)
	}

	if len(contacts) == 0 {
		fmt.Println("No contacts found for", companyName)
		return
	}

	fmt.Printf("\nFound %d contacts for %s:\n", len(contacts), companyName)
	fmt.Printf("Company URL: https://app.apollo.io/#/organizations/%s\n\n", contacts[0].CompanyID)

	for _, c := range contacts {
		fmt.Println("---")
		fmt.Println(c.Name)
		fmt.Println(c.Title)
		fmt.Println(c.ApolloURL)
	}
}
