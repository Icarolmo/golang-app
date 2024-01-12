package eventendpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func CreateJury(c *gin.Context){
	var juriRequest request.JuriRequest
	if err := c.ShouldBindJSON(&juriRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	restError := queries.InsertPessoaJuri(juriRequest)
	if restError != nil {
		c.JSON(restError.Code, restError)
		return
	}
	c.JSON(201, juriRequest)
}