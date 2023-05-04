package controllers

import (
	"api_fiber_warehouse/initializers"
	modelsStruc "api_fiber_warehouse/models/struc/table"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Landing(c *fiber.Ctx) error {
	var Branchs []modelsStruc.Branch
	var err error
	sql_statement := "SELECT id_branch, branch_name, IFNULL(address,'') AS address, IFNULL(id_addr_tambon,0) AS id_addr_tambon, IFNULL(id_addr_amphure, 0) AS id_addr_amphure, IFNULL(id_addr_province, 0) AS id_addr_province, IFNULL(branch_tel, '') as branch_tel, IFNULL(branch_mobile, '') as branch_mobile,  is_active_flag  FROM branch where is_active_flag=1"
	list, err := initializers.DB.Query(sql_statement)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Branch
		err = list.Scan(
			&row.IDBranch,
			&row.BranchName,
			&row.Address,
			&row.IDAddrTambon,
			&row.IDAddrAmphure,
			&row.IDAddrProvince,
			&row.BranchTel,
			&row.BranchMobile,
			&row.IsActiveFlag,
		)
		if err != nil {
			log.Fatal(err.Error())
		}
		Branchs = append(Branchs, row)
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	return c.JSON(fiber.Map{"status": "success", "branchs": Branchs})

}
