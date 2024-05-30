package repositories

import (
	"database/sql"
	"project-indekost/structs"
	"project-indekost/utils"
)

func GetDataRental(db *sql.DB) (result []structs.Rental, err error) {
	query := "SELECT id,lodger_id,room_id,date_start FROM rental"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rent := structs.Rental{}
		rows.Scan(&rent.ID, &rent.LodgerID, &rent.RoomID, &rent.DateStart)
		if err != nil {
			return nil, err
		}
		result = append(result, rent)
	}

	return
}
func InsertDataRental(db *sql.DB, rental structs.Rental) (err error) {
	query := "INSERT INTO rental(lodger_id,room_id,date_start)VALUES($1,$2,$3)"
	errs := db.QueryRow(query, rental.LodgerID, rental.RoomID, rental.DateStart)
	return errs.Err()

}
func DeleteDataRental(db *sql.DB, rental structs.Rental) (err error) {
	query := "DELETE FROM rental WHERE id=$1"
	errs := db.QueryRow(query, rental.ID)
	return errs.Err()
}

func GetDueDateTransaction(db *sql.DB) (result []structs.TransactionDueDate, err error) {
	query := `SELECT id,lodger_id,due_date FROM rent_transaction`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rentTrans := structs.TransactionDueDate{}
		rows.Scan(&rentTrans.ID, &rentTrans.LodgerID, &rentTrans.DueDate)
		if err != nil {
			return nil, err
		}
		result = append(result, rentTrans)
	}

	return
}

func InsertTransactionRentData(db *sql.DB, rentTrans structs.Transaction) (err error) {
	querySql := `INSERT INTO rent_transaction
		(lodger_id, room_id, bank, payment_date, due_date, payment_type, main_fee, additional_fee, total_fee)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	errs := db.QueryRow(querySql, rentTrans.LodgerID, rentTrans.RoomID, rentTrans.Bank, rentTrans.PaymentDate, rentTrans.DueDate, rentTrans.PaymentType, rentTrans.MainFee, rentTrans.AdditionalFee,
		utils.SumFee(uint(rentTrans.MainFee), uint(rentTrans.AdditionalFee)))

	return errs.Err()
}

func DeleteTransactionRentData(db *sql.DB, rentTrans structs.Transaction) (err error) {
	query := "DELETE FROM rent_transaction WHERE id=$1"
	errs := db.QueryRow(query, rentTrans.ID)
	return errs.Err()
}
