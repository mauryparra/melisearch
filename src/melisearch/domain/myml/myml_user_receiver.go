package myml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mauryparra/melisearch/src/melisearch/utils/apierrors"
)

const urlUsers = "https://api.mercadolibre.com/users/"

// Get obtiene toda la iformación del usuario a partir del ID
func (user *User) Get() *apierrors.ApiError {
	if user.ID == 0 {
		return &apierrors.ApiError{
			Message: "userID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%d", urlUsers, user.ID)
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

	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
