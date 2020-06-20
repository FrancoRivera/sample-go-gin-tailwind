package repository

import (
	"errors"
	"../models"
	"strings")

type HouseRepository struct{
	//Repository
	//model models.House
}

var houseList = []models.House{
	models.House{ID: 1, Title: "House 1", Content: "House 1 Listing "},
	models.House{ID: 2, Title: "House 2", Content: "House 2 Listing "},
	models.House{ID: 3, Title: "House 3", Content: "House 3 Listing "},
	models.House{ID: 4, Title: "House 4", Content: "House 4 Listing "},
	models.House{ID: 5, Title: "House 5", Content: "House 5 Listing "},
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


// not implemented

func (r *HouseRepository) FindByTitle(title string) ([]models.House, error) {
	matchedList := []models.House;
	//r.Repository.FetchAll();
	for _, o := range houseList{
		if strings.Contains(o.Title, title){
			matchedList = append(matchedList, o)
		}
	}
	return matchedList
}

