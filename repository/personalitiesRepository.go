package repository

import (
	"apiRest/models"
)

func FindAll() []models.Personality {
	DB := ConnectBd()
	var p = []models.Personality{}
	DB.Find(&p)
	defer CloseConnectionBd(DB)
	return p

}

func FindId(id int) models.Personality {
	DB := ConnectBd()
	var p = models.Personality{Id: id}
	DB.Find(&p)
	defer CloseConnectionBd(DB)
	return p
}

func Create(p *models.Personality) models.Personality {
	DB := ConnectBd()
	DB.Create(&p)
	defer CloseConnectionBd(DB)
	return *p
}

func Delete(id int) models.Personality {
	DB := ConnectBd()
	var p = models.Personality{Id: id}
	DB.Delete(&p)
	defer CloseConnectionBd(DB)
	return p
}

func Update(p *models.Personality) models.Personality {
	DB := ConnectBd()
	DB.Updates(&p)
	defer CloseConnectionBd(DB)
	return *p
}
