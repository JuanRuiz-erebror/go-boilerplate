package repository

import "goprueba/Services/sample/model"

// BottlesAll example
func BottlesAll() ([]model.Bottle, error) {
	return bottles, nil
}

// BottleOne example
func BottleOne(id int) (*model.Bottle, error) {
	for _, v := range bottles {
		if id == v.ID {
			return &v, nil
		}
	}
	return nil, ErrNoRow
}

var bottles = []model.Bottle{
	{ID: 1, Name: "bottle_1", Account: model.Account{ID: 1, Name: "accout_1"}},
	{ID: 2, Name: "bottle_2", Account: model.Account{ID: 2, Name: "accout_2"}},
	{ID: 3, Name: "bottle_3", Account: model.Account{ID: 3, Name: "accout_3"}},
}
