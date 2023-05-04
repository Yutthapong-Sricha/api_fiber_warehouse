package models

import (
	"log"
	"strconv"

	"api_fiber_warehouse/helpers"
	"api_fiber_warehouse/initializers"

	_ "github.com/go-sql-driver/mysql"

	modelsStrucTable "api_fiber_warehouse/models/struc/table"
)

func ListPosition(act string) []modelsStrucTable.StaffPosition {
	var Positions []modelsStrucTable.StaffPosition
	var err error
	sql_statement := "SELECT id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag FROM staff_position "
	if act == "/active_only" {
		sql_statement = sql_statement + " where is_active_flag=1 "
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStrucTable.StaffPosition
		err = list.Scan(
			&row.IDStaffPosition,
			&row.PositionName,
			&row.RowOrderPosition,
			&row.IsActiveFlag,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.IDStaffPosition))

			row.Id_position_enc = id_enc

		}
		Positions = append(Positions, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Positions
}

func GetPosition(id string) modelsStrucTable.StaffPosition {
	var Position modelsStrucTable.StaffPosition
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag FROM staff_position WHERE id_staff_position=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	rows := statement.QueryRow(id_int)
	err = rows.Scan(
		&Position.IDStaffPosition,
		&Position.PositionName,
		&Position.RowOrderPosition,
		&Position.IsActiveFlag,
	)
	if err != nil {
		return Position
	} else {
		id_enc := helpers.Encrypt(strconv.Itoa(Position.IDStaffPosition))

		Position.Id_position_enc = id_enc

	}
	return Position
}
