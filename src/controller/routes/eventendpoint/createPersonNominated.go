package eventendpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func CreatePersonNominated(c *gin.Context){
	var nominadaRequest request.PessoaNomiRequest
	if err := c.ShouldBindJSON(&nominadaRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	restError := queries.InsertEventoPessoaNominada(nominadaRequest)
	if restError != nil {
		c.JSON(restError.Code, restError)
		return
	}
	c.JSON(201, nominadaRequest)

}