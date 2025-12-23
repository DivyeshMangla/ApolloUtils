package main

import (
	"ApolloUtils/internal/apollo"
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

	client := apollo.New(os.Getenv("APOLLO_API_KEY"))

	org, err := client.FindCompany(companyName)
	if err != nil {
		log.Fatal("Error searching for company:", err)
	}

	if org == nil {
		fmt.Printf("No company found for '%s'\n", companyName)
		return
	}

	fmt.Printf("\nCompany found:\n")
	fmt.Printf("Name: %s\n", org.Name)
	fmt.Printf("Apollo URL: https://app.apollo.io/#/organizations/%s\n", org.ID)
}
