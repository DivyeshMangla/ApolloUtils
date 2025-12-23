package apollo

// FindSalesAndMarketing searches for people with sales, marketing, or executive titles
// at a specific organization. Only returns contacts with both email and phone data.
func (c *Client) FindSalesAndMarketing(orgID string) ([]Person, error) {
	payload := map[string]any{
		"organization_ids": []string{orgID},
		"person_titles": []string{
			"Marketing",
			"Marketing Manager",
			"Marketing Director",
			"Marketing Head",
			"Growth",
			"Partnership",
			"Partnerships",
			"Sponsorship",
			"Brand",
			"Brand Manager",
			"Business Development",
			"BD",
			"Sales",
			"Sales Manager",
			"Sales Director",
			"Account Executive",
			"Revenue",
			"Strategic Partnerships",
			"Community",
			"Creator Partnerships",
			"Influencer Marketing",
			"Alliances",
			"Manager",
			"Director",
			"Head",
			"VP",
			"Vice President",
		},
		"person_locations": []string{"India"},
		"has_email":        true,
		"has_phone":        true,
		"page":             1,
		"per_page":         50,
	}

	var resp PeopleSearchResponse
	err := c.post("/api/v1/mixed_people/api_search", payload, &resp)
	if err != nil {
		return nil, err
	}

	// Filter to ensure contacts actually have phone numbers
	var filteredPeople []Person
	for _, person := range resp.People {
		// Check if person has direct phone (Yes/Maybe are acceptable)
		hasPhone := person.HasDirectPhone == "Yes" || person.HasDirectPhone == "Maybe: please request direct dial via people/bulk_match"
		if hasPhone || (len(person.PhoneNumbers) > 0 && person.PhoneNumbers[0].Sanitized != "") {
			filteredPeople = append(filteredPeople, person)
		}
	}

	// Limit to 25 results after filtering
	if len(filteredPeople) > 25 {
		filteredPeople = filteredPeople[:25]
	}

	return filteredPeople, nil
}
