package myml

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mauryparra/melisearch/src/melisearch/services/myml"
	"github.com/mauryparra/melisearch/src/melisearch/utils/apierrors"
)

const (
	paramUserID = "userID"
)

// GetInfo crea un usuario y obtiene sus datos
func GetInfo(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	mymlData, apiErr := myml.GetInfoFromAPI(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, mymlData)
}
