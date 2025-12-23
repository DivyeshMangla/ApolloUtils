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

	fmt.Print("Enter company name(s) (comma-separated for batch): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	input = strings.TrimSpace(input)

	if input == "" {
		log.Fatal("Company name cannot be empty")
	}

	client := apollo.New(os.Getenv("APOLLO_API_KEY"))

	// Split by comma for batch processing (Microsoft,Google,Apple)
	companies := strings.Split(input, ",")

	fmt.Printf("\nSearching for %d companies...\n\n", len(companies))

	for i, companyName := range companies {
		companyName = strings.TrimSpace(companyName)
		if companyName == "" {
			continue
		}

		org, err := client.FindCompany(companyName)
		if err != nil {
			fmt.Printf("%d. %s - Error: %v\n", i+1, companyName, err)
			continue
		}

		if org == nil {
			fmt.Printf("%d. %s - Not found\n", i+1, companyName)
			continue
		}

		fmt.Printf("%d. %s\n", i+1, org.Name)
		fmt.Printf("   Apollo URL: https://app.apollo.io/#/organizations/%s\n\n", org.ID)
	}
}
