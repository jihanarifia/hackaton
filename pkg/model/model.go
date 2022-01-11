package model

type Dependency struct {
	Name    string
	Healthy bool
}

type HealthCheckResponse struct {
	Name         string       `json:"name"`
	Dependencies []Dependency `json:"dependencies"`
	Healthy      bool         `json:"healthy"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
