package response

type ClienteSuccessResponse struct {
	Success bool `json:"success"`

	Message string `json:"message"`

	Data ClienteResponse `json:"data"`
}
