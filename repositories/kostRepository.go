package repositories

import (
	"database/sql"
	"project-indekost/structs"
)

func GetAllLodgers(db *sql.DB) (result []structs.Lodger, err error) {
	sql := "SELECT id,first_name FROM lodger"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		lodger := structs.Lodger{}
		rows.Scan(&lodger.ID, &lodger.Name)
		result = append(result, lodger)
	}

	return
}

func InsertLodger(db *sql.DB) {

}
