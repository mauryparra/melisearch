package myml

import (
	"sync"

	"github.com/mauryparra/melisearch/src/melisearch/domain/myml"
	"github.com/mauryparra/melisearch/src/melisearch/utils/apierrors"
)

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
			if apiMyML.Site.ID != "" {
				myMLData.Site = apiMyML.Site
				continue
			}

			if apiMyML.PaymentMethods[0].ID != "" {
				myMLData.PaymentMethods = apiMyML.PaymentMethods
				continue
			}
		}
	}()

	wg.Add(2)

	go getSiteFromAPI(myMLData.User.SiteID, c)
	go getPaymentMethodsFromAPI(myMLData.User.SiteID, c)

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
		c <- nil
		return
	}

	var payments myml.PaymentMethods

	if apiErr := payments.Get(siteID); apiErr != nil {
		c <- nil
		return
	}

	c <- &myml.MyML{
		PaymentMethods: payments,
	}

	return
}
