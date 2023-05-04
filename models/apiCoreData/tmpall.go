package models

import (
	"api_fiber_warehouse/initializers"
	modelsStrucTable "api_fiber_warehouse/models/struc/table"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Listtambon(c *fiber.Ctx) []modelsStrucTable.AddrTambon {
	var Tambons []modelsStrucTable.AddrTambon
	var err error
	sql_statement := "SELECT * FROM addr_tambon "

	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStrucTable.AddrTambon
		err = list.Scan(
			&row.IDAddrTambon,
			&row.ZipCode,
			&row.TamNameTh,
			&row.TamNameEn,
			&row.IDAddrAmphure,
			&row.CreatedTime,
			&row.UpdatedTime,
			&row.DeletedTime,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			Tambons = append(Tambons, row)
		}
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Tambons
}

func Listamphure(c *fiber.Ctx) []modelsStrucTable.AddrAmphure {
	var Amphure []modelsStrucTable.AddrAmphure
	var err error
	sql_statement := "SELECT * FROM addr_amphure"

	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStrucTable.AddrAmphure
		err = list.Scan(
			&row.IDAddrAmphure,
			&row.AmpNameTh,
			&row.AmpNameEn,
			&row.IDAddrProvince,
			&row.CreatedTime,
			&row.UpdatedTime,
			&row.DeletedTime,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			Amphure = append(Amphure, row)
		}
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Amphure
}

func Listprovince(c *fiber.Ctx) []modelsStrucTable.AddrProvince {
	var Province []modelsStrucTable.AddrProvince
	var err error
	sql_statement := "SELECT * FROM addr_province"

	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStrucTable.AddrProvince
		err = list.Scan(
			&row.IDAddrProvince,
			&row.ProvNameTh,
			&row.ProvNameEn,
			&row.IDAddrProvince,
			&row.CreateTime,
			&row.UpdateTime,
			&row.DeletedTime,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			Province = append(Province, row)
		}
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Province
}
