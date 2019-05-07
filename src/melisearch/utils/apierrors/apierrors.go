package apierrors

// ApiError es un struct personalizado para errores de apis
type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
