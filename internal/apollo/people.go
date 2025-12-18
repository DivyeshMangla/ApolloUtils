package apollo

// FindSalesAndMarketing searches for people with sales, marketing, or executive titles
// at a specific organization. Only returns contacts with both email and phone data.
func (c *Client) FindSalesAndMarketing(orgID string) ([]Person, error) {
	payload := map[string]any{
		"organization_ids": []string{orgID},
		"person_titles": []string{
			"Marketing",
			"Growth",
			"Partnership",
			"Sponsorship",
			"Brand",
			"Business Development",
			"BD",
			"Sales",
			"Account Executive",
			"Revenue",
			"Strategic Partnerships",
			"Community",
			"Creator Partnerships",
			"Influencer Marketing",
			"Alliances",
		},
		"has_email": true,
		"has_phone": true,
		"page":      1,
		"per_page":  5,
	}

	var resp PeopleSearchResponse
	err := c.post("/api/v1/mixed_people/api_search", payload, &resp)
	if err != nil {
		return nil, err
	}

	return resp.People, nil
}
