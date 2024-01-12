package queries

import (
	"fmt"

	"github.com/icarolmo/golang-app/src/configuration/database/mysql_connection"
	"github.com/icarolmo/golang-app/src/configuration/rest_err"
	"github.com/icarolmo/golang-app/src/controller/model/entity/request"
)

func InsertEvento(eventoRequest request.EventoRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	fmt.Sprintln("Chegou aqui também porra" + eventoRequest.Nome + eventoRequest.Nacionalidade)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_eventos (PK_Nome, tipo, nacionalidade, ano_inicio)" +
	"VALUES ('%s','%s','%s','%d-01-01')", eventoRequest.Nome, eventoRequest.Tipo, eventoRequest.Nacionalidade, eventoRequest.Ano_Inicio))
	if err != nil {
		fmt.Sprintln("Estourando erro aqui")
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
		return restError
	}
	return nil	
}

func InsertEventoEdicao(editionRequest request.EdicaoRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_edicao (PK_ano, localizacao, data, FK_tb_eventos_PK_Nome)" +
	"VALUES ('%d-01-01','%s','%s','%s')", editionRequest.Ano, editionRequest.Localizacao, editionRequest.Data, editionRequest.Nome_Evento))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
		return restError
	}
	return nil	
}

func InsertEventoPremio(premioRequest request.PremioRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_premio (PK_Tipo, Nome, FK_tb_edicao_PK_ano, FK_tb_edicao_tb_eventos_PK_Nome)" +
	"VALUES ('%s','%s','%d-01-01','%s')", premioRequest.Tipo, premioRequest.Nome_Premio, premioRequest.Ano_Edicao, premioRequest.Nome_Evento))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
		return restError
	}
	return nil	
}

func InsertEventoPessoaNominada(nominadaRequest request.PessoaNomiRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO E_Nominado (FK_tb_premio_PK_Tipo, FK_tb_premio_tb_edicao_PK_ano, FK_tb_premio_tb_edicao_tb_eventos_PK_Nome, FK_tb_pessoa_PK_NomeArt, FK_tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal, ganhou)" +
	"VALUES ('%s','%d-01-01','%s','%s', '%d-01-01','%s','%s')",nominadaRequest.Tipo_Premio, nominadaRequest.Ano_Edicao, nominadaRequest.Nome_Evento, nominadaRequest.Nome_Artistico, nominadaRequest.Ano_Producao_Filme, nominadaRequest.Titulo_Original_Filme, nominadaRequest.Ganhou))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
		return restError
	}
	return nil	
}

func InsertEventoFilmeNominado(nominadaRequest request.FilmeNomiRequest) *rest_err.RestErr {
	db := mysql_connection.InitConnection()
	defer mysql_connection.CloseConnection(db)
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tb_filme_nominado (tb_filmes_PK_AnoProducao, FK_tb_filmes_PK_TituloOriginal, FK_tb_premio_PK_Tipo, FK_tb_premio_tb_edicao_PK_ano, FK_tb_premio_tb_edicao_tb_eventos_PK_Nome, ganhou)" +
	"VALUES ('%d-01-01','%s','%s','%d-01-01', '%s', '%s')", nominadaRequest.Ano_Producao, nominadaRequest.Titulo_Original, nominadaRequest.Tipo_Premio, nominadaRequest.Ano_Edicao, nominadaRequest.Nome_Evento, nominadaRequest.Ganhou))
	if err != nil {
		restError := rest_err.NewBadRequestServerError(fmt.Sprintf("Erro ao realizar inserção no banco de dados, error: %s", err))
		return restError
	}
	return nil	
}