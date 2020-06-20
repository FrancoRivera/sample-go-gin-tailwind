package repository

import (
	"errors"
	"../models")

type HouseRepository struct{
	//Repository
	//model models.House
}

var houseList = []models.House{
	models.House{ID: 1, Title: "House 1", Content: "House 1 Listing "},
	models.House{ID: 2, Title: "House 2", Content: "House 2 Listing "},
}


func (r *HouseRepository) FetchAll() []models.House {
	//r.Repository.FetchAll();
	return houseList;
}

func (r *HouseRepository) FetchById(id int) (*models.House, error) {
	//r.Repository.FetchAll();
	for _, o := range houseList{
		if id == o.ID{
			return &o, nil
		}
	}
	return nil, errors.New("House not found")
}
