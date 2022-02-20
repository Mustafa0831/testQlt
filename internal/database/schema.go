package database

import "log"

// BuildSchema ---> create Tables
func (d *Database) BuildSchema() error {
	category, err := d.db.Prepare(`CREATE TABLE IF NOT EXISTS Category ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		category_name VARCHAR(150)
								)`)

	defer category.Close()
	CheckErr(err)
	category.Exec()

	payments, err := d.db.Prepare(`CREATE TABLE IF NOT EXISTS Payment (
		"id"	INTEGER,
	"transaction_name"	TEXT,
	"amount"	NUMERIC,
	"transaction_date"	INTEGER,
	"income"	BOOLEAN,
	"comment"	TEXT,
	"category_id"	INTEGER,
	FOREIGN KEY("category_id") REFERENCES "categories",
	PRIMARY KEY("id" AUTOINCREMENT)
	)`)
	defer payments.Close()
	CheckErr(err)
	payments.Exec()
	return nil
}

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
