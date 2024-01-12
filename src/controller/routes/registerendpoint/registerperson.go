package registerendpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

//Registra pessoa no banco de dados conforme dados no corpo da requisição
func RegisterPerson(c *gin.Context){
	var registerRequest request.RegistrarRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	restError := queries.RegisterPersonInMovie(registerRequest)
	if restError != nil{
		c.JSON(restError.Code, restError)
		return
	}
	c.JSON(200, registerRequest)
}
