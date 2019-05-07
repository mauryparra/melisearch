package myml

// MyML es un struct que contiene info de un usuario, el sitio y la moneda asociada
type MyML struct {
	User           User           `json:"user"`
	Site           Site           `json:"site"`
	PaymentMethods PaymentMethods `json:"payment_methods"`
}
