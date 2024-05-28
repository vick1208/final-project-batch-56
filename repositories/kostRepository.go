package repositories

import (
	"database/sql"
	"fmt"
	"project-indekost/structs"
)

func GetAllLodgers(db *sql.DB) (result []structs.Lodger, err error) {
	query := "SELECT id,first_name FROM lodgers"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		lodger := structs.Lodger{}
		err = rows.Scan(&lodger.ID, &lodger.FirstName)
		if err != nil {
			panic(err)
		}
		result = append(result, lodger)
	}
	return
}
