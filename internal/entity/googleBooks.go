package entity

// GoogleBooksError structure to mapper from google books API
type GoogleBooksError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

// GoogleBooksResponse structure to mapper from google books API
type GoogleBooksResponse struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		VolumeInfo struct {
			Title     string   `json:"title"`
			Authors   []string `json:"authors"`
			Publisher string   `json:"publisher"`
		} `json:"volumeInfo"`
	}
}
