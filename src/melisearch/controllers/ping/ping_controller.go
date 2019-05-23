package ping

import (
	"errors"
	"math/rand"

	"github.com/gin-gonic/gin"
)

// Ping answers with string "pong"
func Ping(context *gin.Context) {

	opcion := rand.Intn(2)

	switch opcion {
	case 0:
		context.AbortWithError(500, errors.New("Error en api final"))
	case 1:
		context.String(200, "Pong API")
	}

}
