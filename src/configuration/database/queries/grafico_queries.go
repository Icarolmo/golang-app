package queries

import (
	"fmt"

	"github.com/icarolmo/golang-app/src/configuration/database/mysql_connection"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func SelectActorMostAwards() ([]request.Ator, *rest_err.RestErr) {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)

	rows, err := db.Query(
	"SELECT FK_tb_pessoa_PK_NomeArt, COUNT(*) AS quantidade_ocorrencias " +
	"FROM E_Nominado " + 
	"WHERE ganhou='sim' " +
	"AND FK_tb_pessoa_PK_NomeArt IN ( " +
			"SELECT FK_tb_pessoa_PK_NomeArt " +
			"FROM tb_autor ) " +
	"GROUP BY FK_tb_pessoa_PK_NomeArt " +
	"ORDER BY quantidade_ocorrencias DESC " +
	"LIMIT 10;")
	
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar consulta, error: %s", err))
		return nil, restError
	}

	defer rows.Close()

	var lista_atores []request.Ator
	for rows.Next() {
		var ator request.Ator
		err := rows.Scan(&ator.Nome_Ator, &ator.Numero_Premios)
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar consulta, error: %s", err))
			return nil, restError
		}
		lista_atores = append(lista_atores, ator)
	}
	
	return lista_atores, nil
}

func SelectMoviesMostMoney() ([]request.Ator, *rest_err.RestErr) {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
 	
	rows, err := db.Query(
	"SELECT PK_TituloOriginal, arrecadacao_prim_ano " +
	"FROM tb_filmes " + 
	"ORDER BY arrecadacao_prim_ano DESC " +
	"LIMIT 10;")
	
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar consulta, error: %s", err))
		return nil, restError
	}

	defer rows.Close()

	var lista_atores []request.Ator
	for rows.Next() {
		var ator request.Ator
		err := rows.Scan(&ator.Nome_Ator, &ator.Numero_Premios)
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar consulta, error: %s", err))
			return nil, restError
		}
		lista_atores = append(lista_atores, ator)
	}
	
	return lista_atores, nil
}

func SelectMoviesMostAwards() ([]request.Ator, *rest_err.RestErr) {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
 	
	rows, err := db.Query(
	"SELECT FK_tb_filmes_PK_TituloOriginal, COUNT(*) AS quantidade_ocorrencias " +
	"FROM tb_filme_nominado " + 
	"GROUP BY FK_tb_filmes_PK_TituloOriginal " +
	"ORDER BY quantidade_ocorrencias DESC " +
	"LIMIT 10;")
	
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar consulta, error: %s", err))
		return nil, restError
	}

	defer rows.Close()

	var lista_atores []request.Ator
	for rows.Next() {
		var ator request.Ator
		err := rows.Scan(&ator.Nome_Ator, &ator.Numero_Premios)
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar consulta, error: %s", err))
			return nil, restError
		}
		lista_atores = append(lista_atores, ator)
	}
	
	return lista_atores, nil
}