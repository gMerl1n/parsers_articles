package handlers

type CategoryRequest struct {
	Name         string `json: Name`
	ProviderSign string `json: ProviderSign`
	URL          string `json: Url`
}

type UserRequest struct {
	Name    string
	Surname string
	Age     int
	Email   string
}
