package request

type FilmeRequest struct {
	Ano_Producao		int	`json:"ano_producao"`
	Titulo_Original		string `json:"titulo_original"`
	Classe				string `json:"classe"`
	Idioma_Original		string `json:"idioma_original"`
	Arrecadacao_PrimAno	int	`json:"arrecadacao_prim_ano"`
	Titulos_Brasil		string `json:"titulos_no_brasil"`
	Data_Estreia		string `json:"data_estreia"`
}