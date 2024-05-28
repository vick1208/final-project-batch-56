package repositories

import (
	"database/sql"
	"project-indekost/model"
)

func GetLodgers(db *sql.DB) (resultList []model.Lodger, err error) {
	sqlQuery := "SELECT * FROM lodgers"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		lodgers := model.Lodger{}
		err = rows.Scan(&lodgers.ID, &lodgers.FirstName, &lodgers.LastName, &lodgers.City, &lodgers.Phone)
		if err != nil {
			panic(err)
		}
		resultList = append(resultList, lodgers)
	}
	return
}
