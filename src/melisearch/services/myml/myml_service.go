package myml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/mauryparra/melisearch/src/melisearch/domain/myml"
	"github.com/mauryparra/melisearch/src/melisearch/utils/apierrors"
)

// const urlCategories = "https://api.mercadolibre.com/sites/%s/categories"

// GetInfoFromAPI obtiene la información del usuario, sitio y moneda desde un userID
func GetInfoFromAPI(userID int64) (*myml.MyML, *apierrors.ApiError) {

	var myMLData myml.MyML

	myMLData.User = myml.User{
		ID: userID,
	}
	if apiErr := myMLData.User.Get(); apiErr != nil {
		return nil, apiErr
	}

	c := make(chan *myml.MyML)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 2; i++ {
			apiMyML := <-c
			wg.Done()
			if &apiMyML.Site != nil {
				myMLData.Site = apiMyML.Site
				continue
			}

			if &apiMyML.PaymentMethods != nil {
				myMLData.PaymentMethods = apiMyML.PaymentMethods
				continue
			}
		}
	}()

	wg.Add(2)

	go getSiteFromAPI(myMLData.User.SiteID, c)
	// go getPaymentMethodsFromAPI(myMLData.User.SiteID, c)

	wg.Wait()

	return &myMLData, nil
}

func getSiteFromAPI(siteID string, c chan *myml.MyML) {
	site := myml.Site{
		ID: siteID,
	}
	if apiErr := site.Get(); apiErr != nil {
		c <- nil
		return
	}

	c <- &myml.MyML{
		Site: site,
	}
	return
}

func getPaymentMethodsFromAPI(siteID string, c chan *myml.MyML) {
	if siteID == "" {
		return nil, &apierrors.ApiError{
			Message: "siteID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	var categories []*myml.Category

	final := fmt.Sprintf(urlCategories, siteID)
	response, err := http.Get(final)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &categories); err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}
