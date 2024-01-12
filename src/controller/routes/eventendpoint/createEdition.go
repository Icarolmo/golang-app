package eventendpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func CreateEdition(c *gin.Context) {
	var editionRequest request.EdicaoRequest
	if err := c.ShouldBindJSON(&editionRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	restError := queries.InsertEventoEdicao(editionRequest)
	if restError != nil {
		c.JSON(restError.Code, restError)
		return
	}
	c.JSON(201, editionRequest)
}