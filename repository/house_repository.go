package repository

import (
	"../models"
	"errors"
	"strings"
	"fmt"

  "database/sql"
)

type HouseRepository struct {
	//Repository
	//model models.House
}

func buildHouse(rows *sql.Rows) []models.House{
	houseList := []models.House{}

	defer rows.Close()
    for rows.Next(){
      var id int
      var firstName string
      var lastName string
      var lastUpdate string
      err := rows.Scan(&id, &firstName, &lastName, &lastUpdate)
      if err!= nil{
        panic(err)
      }

      // fmt.Println(id, firstName, lastName, lastUpdate)
      house := models.House{
      	 		ID: id,
      			Title: firstName,
      			Content: lastName,
      		}
      houseList = append(houseList, house)
    }

    err := rows.Err()
    if err != nil {
      panic(err)
    } 
    return houseList
}

func (r *HouseRepository) FetchAll() []models.House {
	rows := QueryDB(`SELECT * FROM public.actor`)
	houseList := buildHouse(rows)

	return houseList
}

func (r *HouseRepository) FetchById(id int) (*models.House, error) {

	sqlStatement := fmt.Sprintf(`SELECT * FROM public.actor WHERE actor_id=%d`, id)
	rows := QueryDB(sqlStatement)
	houseList := buildHouse(rows)

	if len(houseList) > 0{
		return &houseList[0], nil
	} else{
		return nil, errors.New("House not found")
	}
}

// not implemented

func (r *HouseRepository) FindByTitle(title string) []models.House {
	matchedList := []models.House{}

	for _, o := range r.FetchAll(){
		if strings.Contains(o.Title, title) {
			matchedList = append(matchedList, o)
		}
	}
	return matchedList
}
