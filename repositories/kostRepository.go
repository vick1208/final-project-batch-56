package repositories

import (
	"database/sql"
	"project-indekost/structs"
)

func GetAllLodgers(db *sql.DB) (result []structs.Lodger, err error) {
	query := "SELECT id,name,city,phone FROM lodger"
	rows, err := db.Query(query)
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

func UpdateLodger(db *sql.DB, lodger structs.Lodger) error {
	query := "UPDATE lodger SET name=$1,city=$2,phone=$3 WHERE id=$4"
	errs := db.QueryRow(query, lodger.Name, lodger.City, lodger.Phone, lodger.ID)
	return errs.Err()
}
func DeleteLodger(db *sql.DB, lodger structs.Lodger) error {
	query := "DELETE FROM lodger WHERE id=$1"
	errs := db.QueryRow(query, lodger.ID)
	return errs.Err()
}

func GetAllRooms(db *sql.DB) ([]structs.Room, error) {
	var (
		result []structs.Room
		err    error
	)

	query := "SELECT id,room_number,price,room_status FROM room"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		rooms := structs.Room{}
		rows.Scan(&rooms.ID, &rooms.RoomNumber, &rooms.Price, &rooms.RoomStatus)
		result = append(result, rooms)
	}
	return result, err

}

func UpdateRoom(db *sql.DB, room structs.Room) error {
	query := "UPDATE room SET room_number=$1,price=$2,room_status=$3 FROM id=$4"
	errs := db.QueryRow(query, room.RoomNumber, room.Price, room.RoomStatus, room.ID)
	return errs.Err()
}
