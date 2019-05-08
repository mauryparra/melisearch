package myml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mauryparra/melisearch/src/melisearch/utils/apierrors"
)

const urlPayments = "http://localhost:8081/sites/%s/payment_methods"

// Get obtiene toda la iformación de los metodos de pago a partir del site ID
func (payments *PaymentMethods) Get(siteID string) *apierrors.ApiError {

	final := fmt.Sprintf(urlPayments, siteID)

	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), payments); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
