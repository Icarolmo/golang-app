package request

type RegistrarRequest struct {
	Nome_Artistico 	string `json:"nome_artistico"`
	Cargo 			string `json:"cargo"`
	Cargo_Principal bool   `json:"cargo_principal"`
	Titulo_Original string `json:"titulo_original_filme"`
	Ano_Producao 	int `json:"ano_producao_filme"`
	Genero 			string `json:"genero"`
}