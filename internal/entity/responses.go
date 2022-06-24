package entity

type FailureResponse struct {
	Message string `json:"message" example:"Message error for users"`
	Details string `json:"details" example:"Message error for developers"`
	Status  string `json:"status" example:"failure"`
}

// SuccessResponse is the basic form of an informational reply upon success
type SuccessResponse struct {
	Success bool   `json:"success" description:"response status" example:"true"`
	Message string `json:"message" description:"response information" example:"ok"`
	Details string `json:"details,omitempty" description:"optional details" example:"operation completed"`
}
