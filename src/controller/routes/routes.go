package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/controller/routes/eventendpoint"
	"github.com/icarolmo/golang-app/src/controller/routes/graphicsendpoint"
	"github.com/icarolmo/golang-app/src/controller/routes/movieendpoint"
	"github.com/icarolmo/golang-app/src/controller/routes/registerendpoint"
	"github.com/icarolmo/golang-app/src/controller/routes/userendpoint"
)

func InitRoutes(r *gin.RouterGroup) {

	// Pacote que contém as rotas da api

	// Página com gráficos
	r.GET("/graphic", graphicsendpoint.GraphicsPage)

	// Endpoint para gerar dados para gráficos
	r.GET("/atores", graphicsendpoint.GetActorData)
	r.GET("/filmes", graphicsendpoint.GetMovieAwardsData)
	r.GET("/filmes/arrecadacao", graphicsendpoint.GetMovieMoneyData)


	//Registrar Filmes no banco de dados
	r.POST("/movie/registerMovie", movieendpoint.CreateMovie)
	//Registrar Relação entre Pessoas em Filmes
	r.POST("/movie/registerPerson", registerendpoint.RegisterPerson)

	//Registrar Pessoas no banco de dados
	r.POST("/person/registerPerson", userendpoint.CreateUser)
	
	//Registrar Eventos no banco de dados
	r.POST("/event/registerEvent", eventendpoint.CreateEvent)
	// Registrar Edição
	r.POST("/event/registerEdition", eventendpoint.CreateEdition)
	// Registrar Juri
	r.POST("/event/registerJury", eventendpoint.CreateJury)
	// Registrar Prêmio
	r.POST("/event/registerAward", eventendpoint.CreateAward)
	// Registrar pessoa indicada a um prêmio de uma edição
	r.POST("/event/registerPersonNominated", eventendpoint.CreatePersonNominated)
	// Registrar filme indicado a um prêmio de uma edição
	r.POST("/event/registerMovieNominated", eventendpoint.CreateMovieNominated)
	
}