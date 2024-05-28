package repositories

import (
	"database/sql"
	"project-indekost/structs"
)

func GetAllLodgers(db *sql.DB) (results []structs.Lodger, err error) {
	query := "SELECT id,name FROM lodger"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		lodger := structs.Lodger{}
		err = rows.Scan(&lodger.ID, &lodger.Name)
		if err != nil {
			panic(err)
		}
		results = append(results, lodger)
	}

	return
}
