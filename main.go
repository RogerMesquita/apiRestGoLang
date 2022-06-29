package main

import (
	"apiRest/models"
	"apiRest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{
			Id:      1,
			Name:    "Roger Mesquita",
			History: "Tentando ser Programador",
		},
		{
			Id:      2,
			Name:    "Teste",
			History: "demais",
		},
	}

	routes.HandleRequest()
}
