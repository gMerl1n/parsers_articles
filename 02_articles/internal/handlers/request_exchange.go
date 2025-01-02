package handlers

type CategoryRequest struct {
	Name         string `json: Name`
	ProviderSign string `json: ProviderSign`
	URL          string `json: Url`
}
