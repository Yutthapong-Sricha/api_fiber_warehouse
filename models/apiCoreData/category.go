package models

import (
	"api_fiber_warehouse/helpers"
	"api_fiber_warehouse/initializers"
	"log"

	"strconv"

	modelsStrucTable "api_fiber_warehouse/models/struc/table"

	_ "github.com/go-sql-driver/mysql"
)

func ListCategory(act string) []modelsStrucTable.ProdCategory {
	var Categorys []modelsStrucTable.ProdCategory
	var err error
	sql_statement := "SELECT id_prod_category, category_name, is_active_flag, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name FROM prod_category "
	if act == "/active_only" {
		sql_statement = sql_statement + " where is_active_flag=1 "
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStrucTable.ProdCategory
		err = list.Scan(
			&row.IDProdCategory,
			&row.CategoryName,
			&row.IsActiveFlag,
			&row.RecordUpdateTime,
			&row.RecordUpdateByName,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.IDProdCategory))

			row.Id_prod_category_enc = id_enc

		}
		Categorys = append(Categorys, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Categorys
}

func GetCategory(id string) modelsStrucTable.ProdCategory {
	var Category modelsStrucTable.ProdCategory
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT id_prod_category, category_name, is_active_flag, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name FROM prod_category WHERE id_prod_category=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	rows := statement.QueryRow(id_int)
	err = rows.Scan(
		&Category.IDProdCategory,
		&Category.CategoryName,
		&Category.IsActiveFlag,
		&Category.RecordUpdateTime,
		&Category.RecordUpdateByName,
	)
	if err != nil {
		return Category
	} else {
		id_enc := helpers.Encrypt(strconv.Itoa(Category.IDProdCategory))

		Category.Id_prod_category_enc = id_enc

	}
	return Category
}
