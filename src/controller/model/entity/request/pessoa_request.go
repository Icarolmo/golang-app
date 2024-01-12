package request

type PessoaRequest struct {
	Nome_Artistico 	string `json:"nome_artistico"` 
	Nome_Verdadeiro string `json:"nome_verdadeiro"`
	Sexo 	 		string `json:"sexo"`
	Ano_Nascimento 	int	`json:"ano_nascimento"`
	Site			string `json:"site"`
	Ano_Inicio		int	`json:"ano_inicio"`
	Situacao		string `json:"situacao"`
	Nro_Total_Anos	int	`json:"nro_total_anos"`
	Ator 			bool `json:"e_ator"`
	Roteirista 		bool `json:"e_roteirista"`
	Produtor 		bool `json:"e_produtor"`
	Diretor 		bool `json:"e_diretor"`
	Diretor_Principal string `json:"diretor_principal"`
	
}

type JuriRequest struct {
	Nome_Artistico 	string `json:"nome_artistico"`
	Nome_Evento		string `json:"nome_evento"`
	Ano_Edicao		int `json:"ano_edicao"`
}