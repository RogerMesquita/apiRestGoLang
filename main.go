package main

import (
	"apiRest/models"
	"apiRest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{
			Name:    "Roger Mesquita",
			History: "Tentando ser Programador",
		},
		{
			Name:    "Teste",
			History: "demais",
		},
	}

	routes.HandleRequest()
}
