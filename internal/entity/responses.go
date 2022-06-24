package entity

type FailureResponse struct {
	Message string `json:"message" example:"Message error for users"`
	Details string `json:"details" example:"Message error for developers"`
	Status  string `json:"status" example:"failure"`
}
