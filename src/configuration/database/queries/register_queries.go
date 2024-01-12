package queries

import (
	"fmt"
	"strings"

	"github.com/icarolmo/golang-app/src/configuration/database/mysql_connection"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func RegisterPersonInMovie(registerRequest request.RegistrarRequest) *rest_err.RestErr{

	switch strings.ToUpper(registerRequest.Cargo) {
    case "ATOR":
		
		db := mysql_connection.InitConnection()
		defer mysql_connection.CloseConnection(db)
		if registerRequest.Cargo_Principal {
			_, err := db.Exec(fmt.Sprintf("INSERT INTO Ator_Principal (FK_tb_autor_tb_pessoa_PK_NomeArt, FK_tb_outros_tb_filmes_PK_AnoProducao, FK_tb_outros_tb_filmes_PK_TituloOriginal)" +
			"VALUES ('%s', '%d-01-01', '%s')", registerRequest.Nome_Artistico, registerRequest.Ano_Producao, registerRequest.Titulo_Original))
			if err != nil {
				restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
				return restError
			}
		} else {
			_, err := db.Exec(fmt.Sprintf("INSERT INTO Ator_Elenco (FK_tb_autor_tb_pessoa_PK_NomeArt, FK_tb_outros_tb_filmes_PK_AnoProducao, FK_tb_outros_tb_filmes_PK_TituloOriginal)" +
			"VALUES ('%s', '%d-01-01', '%s')", registerRequest.Nome_Artistico, registerRequest.Ano_Producao, registerRequest.Titulo_Original))
			if err != nil {
				restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
				return restError
			}
		}
		return nil
    case "ROTEIRISTA":
		
		restError := SelectRoteirista(registerRequest.Nome_Artistico)
		if restError != nil {
			return restError
		}
		if strings.ToUpper(registerRequest.Genero) == "DOCUMENTARIO" {
			restError2 := SelectMovieDocumentario(registerRequest.Ano_Producao, registerRequest.Titulo_Original)
			if restError2 != nil {
				return restError2
			}
		} else {
			restError2 := SelectMovieOutros(registerRequest.Ano_Producao, registerRequest.Titulo_Original)
			if restError2 != nil {
				return restError2
			}
		}

		db := mysql_connection.InitConnection()
		defer mysql_connection.CloseConnection(db)
		_, err := db.Exec(fmt.Sprintf("INSERT INTO E_Roteirista (FK_tb_roteirista_tb_pessoa_PK_NomeArt, FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal)" +
		"VALUES ('%s', '%d-01-01', '%s')", registerRequest.Nome_Artistico, registerRequest.Ano_Producao, registerRequest.Titulo_Original))
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
			return restError
		}
		return nil
    case "PRODUTOR":
		restError := SelectProdutor(registerRequest.Nome_Artistico)
		if restError != nil {
			return restError
		} 
		if strings.ToUpper(registerRequest.Genero) == "DOCUMENTARIO" {
			restError2 := SelectMovieDocumentario(registerRequest.Ano_Producao, registerRequest.Titulo_Original)
			if restError2 != nil {
				return restError2
			}
		} else {
			restError2 := SelectMovieOutros(registerRequest.Ano_Producao, registerRequest.Titulo_Original)
			if restError2 != nil {
				return restError2
			}
		}

		db := mysql_connection.InitConnection()
		defer mysql_connection.CloseConnection(db)
		_, err := db.Exec(fmt.Sprintf("INSERT INTO E_Produtor (FK_tb_produtor_tb_pessoa_PK_NomeArt, FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal)" +
		"VALUES ('%s', '%d-01-01', '%s')", registerRequest.Nome_Artistico, registerRequest.Ano_Producao, registerRequest.Titulo_Original))
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
			return restError
		}
		return nil
    case "DIRETOR":
		
		restError := SelectDiretor(registerRequest.Nome_Artistico)
		if restError != nil {
			return restError
		}
		db := mysql_connection.InitConnection()
		defer mysql_connection.CloseConnection(db)
		_, err := db.Exec(fmt.Sprintf("INSERT INTO E_Diretor (FK_tb_diretor_tb_pessoa_PK_NomeArt, FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal)" +
		"VALUES ('%s', '%d-01-01', '%s')", registerRequest.Nome_Artistico, registerRequest.Ano_Producao, registerRequest.Titulo_Original))
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
			return restError
		}
		return nil
    default:
        restErr := rest_err.NewBadRequestError(fmt.Sprintf("Erro no campo cargo, valor: %s inválido", registerRequest.Cargo))
		return restErr
    }
}