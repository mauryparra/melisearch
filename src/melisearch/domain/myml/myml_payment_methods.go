package myml

// PaymentMethods es la representación de conjunto de metodos de pago de por mercado libre
type PaymentMethods []struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeID   string `json:"payment_type_id"`
	Thumbnail       string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}
