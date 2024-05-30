package repositories

import (
	"database/sql"
	"project-indekost/structs"
)

func GetAllLodgers(db *sql.DB) (result []structs.Lodger, err error) {
	query := "SELECT id,name,city,phone FROM lodger ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		lodger := structs.Lodger{}
		err = rows.Scan(&lodger.ID, &lodger.Name, &lodger.City, &lodger.Phone)
		if err != nil {
			return
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

func GetAllRooms(db *sql.DB) (result []structs.Room, err error) {

	query := "SELECT id,room_number,price,room_status FROM room ORDER BY id"
	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		rooms := structs.Room{}
		rows.Scan(&rooms.ID, &rooms.RoomNumber, &rooms.Price, &rooms.RoomStatus)
		result = append(result, rooms)
	}
	return result, err

}

func InsertRoom(db *sql.DB, room structs.Room) error {
	query := "INSERT INTO room (room_number,price,room_status) VALUES ($1,$2,$3)"
	errs := db.QueryRow(query, room.RoomNumber, room.Price, room.RoomStatus)
	return errs.Err()
}
func UpdateRoom(db *sql.DB, room structs.Room) error {
	query := "UPDATE room SET room_number=$1,price=$2,room_status=$3 WHERE id=$4"
	errs := db.QueryRow(query, room.RoomNumber, room.Price, room.RoomStatus, room.ID)
	return errs.Err()
}

func DeleteRoom(db *sql.DB, room structs.Room) (err error) {
	query := "DELETE FROM room WHERE id=$1"
	errs := db.QueryRow(query, room.ID)
	return errs.Err()
}
