package apollo

// CompanySearchResponse contains organization search results from Apollo API.
type CompanySearchResponse struct {
	Organizations []Organization `json:"organizations"`
}

// Organization represents a company entity in Apollo.
type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PeopleSearchResponse contains person search results from Apollo API.
type PeopleSearchResponse struct {
	People []Person `json:"people"`
}

// Person represents an individual contact in Apollo with their professional details.
type Person struct {
	ID             string        `json:"id"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	Title          string        `json:"title"`
	Email          string        `json:"email"`
	PhoneNumbers   []PhoneNumber `json:"phone_numbers"`
	HasDirectPhone string        `json:"has_direct_phone"`
}

// PhoneNumber represents a contact's phone number with standardized formatting.
type PhoneNumber struct {
	Sanitized string `json:"sanitized_number"`
}
