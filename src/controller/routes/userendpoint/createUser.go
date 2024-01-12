package userendpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/configuration/database/queries"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)



func CreateUser(c *gin.Context){
	var pessoaRequest request.PessoaRequest
	if err := c.ShouldBindJSON(&pessoaRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))
		c.JSON(restErr.Code, restErr)	
		return
	}

	if pessoaRequest.Ator || pessoaRequest.Diretor || pessoaRequest.Produtor || pessoaRequest.Roteirista {
		row, restError := queries.SelectPessoa(pessoaRequest.Nome_Artistico)
		if restError != nil {
			c.JSON(restError.Code, restError)	
			return
		}
		if row != 0 {
			restError := rest_err.NewBadRequestError(fmt.Sprintf("Pessoa com nome_artistico=%s já existente no banco", pessoaRequest.Nome_Artistico))
			c.JSON(restError.Code, restError)
			return
		} else {
			restError := queries.InsertPessoa(pessoaRequest)
			if restError != nil {
				c.JSON(restError.Code, restError)
				return
			}
		}
	} else {
		restError := rest_err.NewBadRequestError("Pessoa deve exercer algum dos cargos: Ator, Diretor, Roteirista ou Produtor")
		c.JSON(restError.Code, restError)
		return
	}

	if pessoaRequest.Ator {
		restError := queries.InsertCargoAtor(pessoaRequest.Nome_Artistico)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
	}

	if pessoaRequest.Roteirista {
		restError := queries.InsertCargoRoteirista(pessoaRequest.Nome_Artistico)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
	}

	if pessoaRequest.Produtor {
		restError := queries.InsertCargoProdutor(pessoaRequest.Nome_Artistico)
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
	}

	if pessoaRequest.Diretor {
		restError := queries.InsertCargoDiretor(pessoaRequest.Nome_Artistico, pessoaRequest.Diretor_Principal )
		if restError != nil {
			c.JSON(restError.Code, restError)
			return
		}
	}
	c.JSON(201, pessoaRequest) 
}