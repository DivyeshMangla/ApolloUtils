package app

import (
	"ApolloUtils/internal/apollo"
	"ApolloUtils/internal/domain"
	"os"
)

// App provides business operations using the Apollo API.
type App struct {
	apollo *apollo.Client
}

// New creates a new App instance with Apollo API credentials from environment.
func New() *App {
	return &App{
		apollo: apollo.New(os.Getenv("APOLLO_API_KEY")),
	}
}

// FindSponsors searches for sales and marketing contacts at a company.
// Returns contacts with names, titles, and Apollo profile URLs.
func (a *App) FindSponsors(company string) ([]domain.Contact, error) {
	org, err := a.apollo.FindCompany(company)
	if err != nil || org == nil {
		return nil, err
	}

	people, err := a.apollo.FindSalesAndMarketing(org.ID)
	if err != nil {
		return nil, err
	}

	contacts := make([]domain.Contact, 0, len(people))

	for _, p := range people {
		contact := domain.Contact{
			Name:      p.FirstName + " " + p.LastName,
			Title:     p.Title,
			ApolloURL: "https://app.apollo.io/#/people/" + p.ID,
			CompanyID: org.ID,
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}
