package eventendpoint

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func CreateAward(c *gin.Context) {
	var premioRequest request.PremioRequest
	if err := c.ShouldBindJSON(&premioRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campo(s) incorreto(s), error=%s", err))
		c.JSON(restErr.Code, restErr)
		return
	}

	switch strings.ToLower(premioRequest.Tipo) {
    case "melhor ator principal":
		restError := queries.InsertEventoPremio(premioRequest)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
		c.JSON(201, premioRequest)	
		return
    case "melhor atriz principal":
		restError := queries.InsertEventoPremio(premioRequest)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
		c.JSON(201, premioRequest)
		return
    case "melhor ator coadjuvante":
		restError := queries.InsertEventoPremio(premioRequest)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
		c.JSON(201, premioRequest)
		return
    case "melhor atriz coadjuvante":
		restError := queries.InsertEventoPremio(premioRequest)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
		c.JSON(201, premioRequest)
		return
	case "melhor filme":
		restError := queries.InsertEventoPremio(premioRequest)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
		c.JSON(201, premioRequest)
		return
	case "melhor direção":
		restError := queries.InsertEventoPremio(premioRequest)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
		c.JSON(201, premioRequest)
		return
    default:
        restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campo tipo_premio=%s incorreto, valores aceitos: 'melhor ator principal', 'melhor atriz principal', " +
		"'melhor ator coadjuvante', 'melhor atriz coadjuvante', 'melhor filme' e 'melhor direção'", premioRequest.Tipo))
		c.JSON(restErr.Code, restErr)	
		return
    }
}