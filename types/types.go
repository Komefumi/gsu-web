package types

type StandardResponse struct {
	Success uint8       `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

type PayloadFieldValidationError struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
