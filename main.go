package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/icarolmo/golang-app/src/controller/routes"
	"github.com/joho/godotenv"
)

func main(){

	// Iniciando configurações defaults de um servidor para subir a aplicação
	router := gin.Default()
	
	// Definindo pasta para busca de templates pelo gin
	router.LoadHTMLGlob("templates/*")

	// Iniciando rotas definidas
	routes.InitRoutes(&router.RouterGroup)

	// Carregando variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Startando servidor na porta definida nas variáveis de ambiente
	if err:= router.Run(":"+os.Getenv("SERVER_PORT")); err != nil {
		log.Fatal(err)
	}
}