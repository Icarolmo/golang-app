package request

type PremioRequest struct {
	Nome_Premio string `json:"nome_premio"`
	Tipo 		string `json:"tipo_premio"`
	Nome_Evento string `json:"nome_evento"`
	Ano_Edicao	int `json:"ano_edicao"`
}