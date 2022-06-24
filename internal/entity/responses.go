package entity

type FailureResponse struct {
	Message string `json:"message" example:"Error message for users"`
	Details string `json:"details" example:"Error message for developers"`
	Status  string `json:"status" example:"failure"`
}
