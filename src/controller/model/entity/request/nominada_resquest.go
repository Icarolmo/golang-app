package request

type PessoaNomiRequest struct {
	Tipo_Premio 		  string `json:"tipo_premio"`
	Ano_Edicao 			  int `json:"ano_edicao"`
	Nome_Evento 		  string `json:"nome_evento"`
	Nome_Artistico 		  string `json:"nome_artistico"`
	Ano_Producao_Filme 	  int `json:"ano_producao_filme"`
	Titulo_Original_Filme string `json:"titulo_original_filme"`
	Ganhou 			      string `json:"ganhou"`
}

type FilmeNomiRequest struct {
	Ano_Producao 	int `json:"ano_producao"`
	Titulo_Original string `json:"titulo_original"`
	Tipo_Premio 	string `json:"tipo_premio"`
	Ano_Edicao 		int `json:"ano_edicao"`
	Nome_Evento 	string `json:"nome_evento"`
	Ganhou 			string `json:"ganhou"`
}