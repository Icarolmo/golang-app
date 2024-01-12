package entity

type Pessoa struct {
	Nome_Artistico 	string 	`json:"nome_artistico"` 
	Nome_Verdadeiro string	`json:"nome_verdadeiro"`
	Sexo 	 		string	`json:"sexo"`
	Ano_Nascimento 	string	`json:"ano_nascimento"`
	Site			string	`json:"site"`
	Ano_Inicio		string	`json:"ano_inicio"`
	Situacao		string	`json:"situacao"`
	Nro_Total_Anos	int		`json:"nro_total_anos"`
}