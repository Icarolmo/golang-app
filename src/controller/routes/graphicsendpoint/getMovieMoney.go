package graphicsendpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
)

func GetMovieMoneyData(c *gin.Context){

	lista_atores, restError := queries.SelectMoviesMostMoney()
	if restError != nil {
		c.JSON(restError.Code, restError)
		return
	}
	c.JSON(200, lista_atores)
}