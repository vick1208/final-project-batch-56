package repositories

import (
	"database/sql"
	"project-indekost/structs"
)

func GetAllLodgers(db *sql.DB) (result []structs.Lodger, err error) {
	sql := "SELECT id,name,city,phone FROM lodger"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		lodger := structs.Lodger{}
		err = rows.Scan(&lodger.ID, &lodger.Name, &lodger.City, &lodger.Phone)
		if err != nil {
			panic(err)
		}
		result = append(result, lodger)
	}

	return
}

func InsertLodger(db *sql.DB, lodger structs.Lodger) error {
	query := "INSERT INTO lodger(name,city,phone) VALUES($1,$2,$3)"
	errs := db.QueryRow(query, lodger.Name, lodger.City, lodger.Phone)
	return errs.Err()
}
