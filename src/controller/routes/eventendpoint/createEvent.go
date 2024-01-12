package eventendpoint

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func CreateEvent(c *gin.Context){
	var eventoRequest request.EventoRequest
	if err := c.ShouldBindJSON(&eventoRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	tipo_evento := strings.ToUpper(eventoRequest.Tipo)

	if tipo_evento !=  "ACADEMIA" && tipo_evento !=  "FESTIVAL" && tipo_evento !=  "CONCURSO" {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campo tipo igual a %s inválido, valores aceitos: 'academia', 'festival' ou 'concurso'", eventoRequest.Tipo))
		c.JSON(restErr.Code, restErr)	
		return
	}
    
	restError := queries.InsertEvento(eventoRequest)
	if restError != nil {
		c.JSON(restError.Code, restError)
		return
	}
	c.JSON(201, eventoRequest)
}
