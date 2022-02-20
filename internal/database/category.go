package database

import (
	"fmt"
	"qlt/internal/model"
)

// Create ---> signing up
func (d *Database) Create(u *model.PaymentDB) (*model.PaymentDB, error) {
	if err := d.Open(); err != nil {
		return nil, err
	}

	stmnt, err := d.db.Prepare("INSERT INTO Category (id,category) VALUES (?,?)")
	res, err := stmnt.Exec(u.Category, u.Payment)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// assign UserID to model 'User'
	id, _ := res.LastInsertId()
	u.ID = id

	return u, nil

}
