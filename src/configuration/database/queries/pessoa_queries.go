package queries

import (
	"fmt"

	"github.com/icarolmo/golang-app/src/configuration/database/mysql_connection"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func SelectPessoa(nome_artistico string) (int,*rest_err.RestErr) {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	rows, err := db.Query("SELECT * FROM tb_pessoa WHERE PK_NomeArt = ?", nome_artistico)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta no banco, error:%s", err))
		return 0, restError
	}
	defer rows.Close()

	if rows.Next() {
		return 1,nil
	} 
	return 0, nil
}

func SelectAtor(nome_artistico string) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	rows, err := db.Query("SELECT * FROM tb_autor WHERE FK_tb_pessoa_PK_NomeArt = ?", nome_artistico)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta no banco, error:%s", err))
		return restError
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}
	restError := rest_err.NewNotFoundError(fmt.Sprintf("Ator %s não existe no banco", nome_artistico))
	return restError
}

func SelectRoteirista(nome_artistico string) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	rows, err := db.Query("SELECT FK_tb_pessoa_PK_NomeArt FROM tb_roteirista WHERE FK_tb_pessoa_PK_NomeArt = ?", nome_artistico)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta no servidor, error: %s", err))
		return restError
	}

	defer rows.Close()
	
	if rows.Next() {
		return nil
	}

	restError := rest_err.NewNotFoundError(fmt.Sprintf("Roteirista %s não existe no banco", nome_artistico))
	return restError
}

func SelectProdutor(nome_artistico string) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	rows, err := db.Query("SELECT FK_tb_pessoa_PK_NomeArt FROM tb_produtor WHERE FK_tb_pessoa_PK_NomeArt = ?", nome_artistico)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta no servidor, error: %s", err))
		return restError
	}

	defer rows.Close()
	if rows.Next() {
		return nil
	}
	restError := rest_err.NewNotFoundError(fmt.Sprintf("Produtor %s não existe no banco", nome_artistico))
	return restError
}

func SelectDiretor(nome_artistico string) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	rows, err := db.Query("SELECT * FROM tb_diretor WHERE FK_tb_pessoa_PK_NomeArt = ?", nome_artistico)
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar consulta no servidor, error: %s", err))
		return restError
	}
	defer rows.Close()
	
	if rows.Next() {
		return nil
	}
	restError := rest_err.NewNotFoundError(fmt.Sprintf("Produtor %s não existe no banco", nome_artistico))
	return restError
}

func InsertPessoa(pessoa request.PessoaRequest) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_pessoa (PK_NomeArt, nome_verdadeiro, sexo, ano_nasc, site, ano_inic, situacao, nro_total_anos)" + 
	"VALUES ('%s','%s', '%s', '%d-01-01', '%s', '%d-01-01', '%s', %d)", pessoa.Nome_Artistico, pessoa.Nome_Verdadeiro, pessoa.Sexo, pessoa.Ano_Nascimento, pessoa.Site, pessoa.Ano_Inicio, pessoa.Situacao, pessoa.Nro_Total_Anos))
	
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar inserção dos dados no banco, error: %s", err))
		return restError
	}
	return nil
}

func InsertCargoAtor(nome_artistico string) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_autor (FK_tb_pessoa_PK_NomeArt) VALUES ('%s')", nome_artistico))

	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar inserção dos dados no banco, error: %s", err))
		return restError
	}
	return nil
}

func InsertCargoRoteirista(nome_artistico string) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_roteirista (FK_tb_pessoa_PK_NomeArt) VALUES ('%s')", nome_artistico))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar inserção dos dados no banco, error: %s", err))
		return restError
	}
	return nil
}

func InsertCargoProdutor(nome_artistico string) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_produtor (FK_tb_pessoa_PK_NomeArt) VALUES ('%s')", nome_artistico))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar inserção dos dados no banco, error: %s", err))
		return restError
	}
	return nil
}

func InsertCargoDiretor(nome_artistico string, principal string) *rest_err.RestErr{
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_diretor (FK_tb_pessoa_PK_NomeArt, e_principal) VALUES ('%s', '%s')", nome_artistico, principal))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao reaizar inserção dos dados no banco, error: %s", err))
		return restError
	}
	return nil
}

func InsertPessoaJuri(juriRequest request.JuriRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO e_juri (FK_tb_edicao_PK_ano, FK_tb_edicao_tb_eventos_PK_Nome, FK_tb_pessoa_PK_NomeArt) " +
	"VALUES ('%d-01-01', '%s', '%s')", juriRequest.Ano_Edicao, juriRequest.Nome_Evento, juriRequest.Nome_Artistico))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro na inserção no banco de dados, error: %s", err))
		return restError
	}
	return nil	
}
	

