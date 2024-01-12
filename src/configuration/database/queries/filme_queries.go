package queries

import (
	"fmt"
	"strconv"

	"github.com/icarolmo/golang-app/src/configuration/database/mysql_connection"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	entity "github.com/icarolmo/golang-app/src/controller/model/entity/filme"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func SelectMovie(ano_producao int, titulo_original string) (entity.Filme, *rest_err.RestErr) {
	db := mysql_connection.InitConnection()
	ano_producao_query := strconv.Itoa(ano_producao)+"-01-01"
	rows, err := db.Query("SELECT PK_AnoProducao, PK_TituloOriginal FROM tb_filmes WHERE PK_AnoProducao = ? AND PK_TituloOriginal = ?", ano_producao_query, titulo_original)
	if err != nil {
		mysql_connection.CloseConnection(db)
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta: %s", err))
		var movie entity.Filme
		return movie, restError
	}
	defer rows.Close()

	if rows.Next() {
		var movie entity.Filme
		err := rows.Scan(&movie.Ano_Producao, &movie.Titulo_Original, &movie.Classe, &movie.Idioma_Original, &movie.Arrecadacao_PrimAno, &movie.Titulos_Brasil, &movie.Data_Estreia)
		if err != nil {
			mysql_connection.CloseConnection(db)
			restError := rest_err.NewBadRequestServerError("Error ao realizar consulta")
			return movie, restError
		}
		mysql_connection.CloseConnection(db)
		var restError rest_err.RestErr
		return movie, &restError
	}
	restError := rest_err.NewNotFoundError(fmt.Sprintf("Filme %s com ano de producao %d não existe no banco", titulo_original, ano_producao))
	var movie entity.Filme
	return movie, restError
}

func SelectMovieDocumentario(ano_producao int, titulo_original string) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	ano_producao_query := strconv.Itoa(ano_producao)+"-01-01"
	rows, err := db.Query("SELECT FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal FROM tb_documentarios WHERE FK_tb_filmes_PK_AnoProducao = ? AND FK_tb_filmes_PK_TituloOriginal = ?", ano_producao_query, titulo_original)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta, error:%s", err))
		return restError
	}
	defer rows.Close()

	if rows.Next() {
		var ano_filme string 
		var titulo_filme string
		err := rows.Scan(&ano_filme, &titulo_filme) 
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta, error:%s", err))
			return restError
		}
		if titulo_filme == "" {
			restError := rest_err.NewNotFoundError(fmt.Sprintf("Filme %s com ano de producao %d não existente no banco com gênero Documentario", titulo_original, ano_producao))
			return restError
		}
		return nil
	}
	restError := rest_err.NewNotFoundError(fmt.Sprintf("Filme %s com ano de producao %d não existente no banco com gênero Documentario", titulo_original, ano_producao))
	return restError
}

func SelectMovieOutros(ano_producao int, titulo_original string) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	ano_producao_query := strconv.Itoa(ano_producao)+"-01-01"
	rows, err := db.Query("SELECT FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal FROM tb_outros WHERE FK_tb_filmes_PK_AnoProducao = ? AND FK_tb_filmes_PK_TituloOriginal = ?", ano_producao_query, titulo_original)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta, error:%s", err))
		return restError
	}
	defer rows.Close()

	if rows.Next() {
		var ano_filme string 
		var titulo_filme string
		err := rows.Scan(&ano_filme, &titulo_filme) 
		if err != nil {
			restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta, error:%s", err))
			return restError
		}
		if titulo_filme == "" {
			restError := rest_err.NewNotFoundError(fmt.Sprintf("Filme %s com ano de producao %d não existente no banco com gênero Outros", titulo_original, ano_producao))
			return restError
		}
		return nil
	}
	restError := rest_err.NewNotFoundError(fmt.Sprintf("Filme %s com ano de producao %d não existente no banco com gênero Outros", titulo_original, ano_producao))
	return restError
}

func InsertFilme(filme request.FilmeRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_filmes (PK_AnoProducao, PK_TituloOriginal, classe, idioma_original, arrecadacao_prim_ano, titulos_no_brasil, data_estreia)" +
	"VALUES('%d-01-01','%s','%s','%s', %d, '%s', '%s')", filme.Ano_Producao, filme.Titulo_Original, filme.Classe, filme.Idioma_Original, filme.Arrecadacao_PrimAno, filme.Titulos_Brasil, filme.Data_Estreia ))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção dos dados no banco, error:%s", err))
		return restError
	}
	mysql_connection.CloseConnection(db)
	return nil	
}

func InsertFilmeLocalEstreia(ano_producao int, titulo_original string) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_local_estreira (FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal) VALUES ('%d-01-01','%s')", ano_producao, titulo_original))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção dos dados no banco, error:%s", err))
		return restError
	}
	return nil	
}

func InsertClasseDocumentario(ano_producao int, titulo_original string) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_documentarios (FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal) VALUES ('%d-01-01','%s')", ano_producao, titulo_original))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção dos dados no banco, error:%s", err))
		return restError
	}
	return nil	
}

func InsertClasseOutros(ano_producao int, titulo_original string) error{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_outros (FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal) VALUES ('%d-01-01','%s')", ano_producao, titulo_original))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção dos dados no banco, error:%s", err))
		return restError
	}
	return nil	
}