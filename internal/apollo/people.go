package apollo

// FindSalesAndMarketing searches for people with sales, marketing, or executive titles
// at a specific organization. Only returns contacts with both email and phone data.
func (c *Client) FindSalesAndMarketing(orgID string) ([]Person, error) {
	payload := map[string]any{
		"organization_ids": []string{orgID},
		"person_titles": []string{
			"Sales",
			"Marketing",
			"Executive",
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
