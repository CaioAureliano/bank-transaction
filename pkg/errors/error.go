package errors

type HttpErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewHttpError(msg string, status int, err string) *HttpErrorResponse {
	return &HttpErrorResponse{
		Message: msg,
		Status:  status,
		Error:   err,
	}
}
