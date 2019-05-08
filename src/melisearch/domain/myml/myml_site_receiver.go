package myml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mauryparra/melisearch/src/melisearch/utils/apierrors"
)

const urlSites = "http://localhost:8081/sites/"

// Get obtiene toda la iformación del usuario a partir del ID
func (site *Site) Get() *apierrors.ApiError {
	if site.ID == "" {
		return &apierrors.ApiError{
			Message: "siteID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s", urlSites, site.ID)
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

	if err := json.Unmarshal([]byte(data), &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
