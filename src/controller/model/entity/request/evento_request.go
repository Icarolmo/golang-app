package request

type EventoRequest struct {
	Nome 		  string `json:"nome"`
	Tipo 		  string `json:"tipo"`
	Nacionalidade string `json:"nacionalidade"`
	Ano_Inicio    int `json:"ano_inicio"`
}

type EdicaoRequest struct {
	Nome_Evento string `json:"nome_evento"`
	Ano 		int `json:"ano"`
	Localizacao string `json:"localizacao"`
	Data 		string `json:"data"`
}