package movieendpoint

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

//Cria filme no banco de dados conforme dados no corpo da requisição
func CreateMovie(c *gin.Context){
	var filmeRequest request.FilmeRequest
	if err := c.ShouldBindJSON(&filmeRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	switch strings.ToUpper(filmeRequest.Classe) {
    case "DOCUMENTARIO":
		queries.InsertFilme(filmeRequest)
		queries.InsertClasseDocumentario(filmeRequest.Ano_Producao, filmeRequest.Titulo_Original)
		queries.InsertFilmeLocalEstreia(filmeRequest.Ano_Producao, filmeRequest.Titulo_Original)
		c.JSON(201, filmeRequest)	
		return
    default:
        queries.InsertFilme(filmeRequest)
		queries.InsertClasseOutros(filmeRequest.Ano_Producao, filmeRequest.Titulo_Original)
		queries.InsertFilmeLocalEstreia(filmeRequest.Ano_Producao, filmeRequest.Titulo_Original)
		c.JSON(201, filmeRequest)	
		return
	}
}