package apollo

import "strings"

// FindCompany searches for a company by name in Apollo's organization database.
// Returns the first exact match or the top result if no exact match exists.
func (c *Client) FindCompany(name string) (*Organization, error) {
	payload := map[string]any{
		"q_organization_name": name,
		"page":                1,
		"per_page":            10,
	}

	var resp CompanySearchResponse
	err := c.post("/api/v1/organizations/search", payload, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Organizations) == 0 {
		return nil, nil
	}

	// Prefer exact name matches
	for _, org := range resp.Organizations {
		if strings.EqualFold(org.Name, name) {
			return &org, nil
		}
	}

	return &resp.Organizations[0], nil
}
