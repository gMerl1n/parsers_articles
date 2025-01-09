package handlers

type CategoryRequest struct {
	Name         string `json:"Name"`
	ProviderSign string `json:"ProviderSign"`
	URL          string `json:"Url"`
}

type UserRequest struct {
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	Age     int    `json:"Age"`
	Email   string `json:"Email"`
}
