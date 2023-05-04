package models

import (
	"api_fiber_warehouse/helpers"
	"api_fiber_warehouse/initializers"
	modelsStruc "api_fiber_warehouse/models/struc/session"
	modelsStrucTable "api_fiber_warehouse/models/struc/table"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Createtmpcore(c *fiber.Ctx) (modelsStrucTable.TmpDocCore, error) {
	var InstTmpDoccore modelsStrucTable.TmpDocCore
	var err error
	now := time.Now()
	year_20yy := now.Year() // type int
	month_mm := int(now.Month())

	if err := c.BodyParser(&InstTmpDoccore); err != nil {
		myError := errors.New("BodyParser")
		return InstTmpDoccore, myError
	} else {
		InstTmpDoccore.GroupType1in2out = 2
		InstTmpDoccore.Year20yy = strconv.Itoa(year_20yy)
		InstTmpDoccore.MonthMm = strconv.Itoa(month_mm)
	}

	id_tmp_doc_core := InstTmpDoccore.IDTmpDocCore

	checkdup, err_chk := initializers.DB.Prepare("SELECT id_tmp_doc_core  FROM tmp_doc_core WHERE id_tmp_doc_core = ?")
	if err_chk != nil {
		log.Fatal(err_chk.Error())
		myError := errors.New("1 : " + err_chk.Error())
		return InstTmpDoccore, myError
	}
	defer checkdup.Close()
	row_dup := checkdup.QueryRow(id_tmp_doc_core)
	var id_chk_dup = 0
	err = row_dup.Scan(
		&id_chk_dup,
	)
	// if err == nil {
	// 	myError := errors.New("extra check: 1 nil")
	// 	return InstTmpDoccore, myError
	// } else {
	// 	myError := errors.New("extra check: 2 not nil" + err.Error())
	// 	return InstTmpDoccore, myError
	// }
	if err == nil { // err ไม่มี แสดงว่า เจอข้อมูล
		// update
		sqlStmt := `UPDATE tmp_doc_core
				SET id_customer = ?,
					tmp_customer_name = ?,
					tmp_customer_mobile = ?,
					id_cust_recomm = ?,
					tmp_cust_name_recomm = ?,
					tmp_cust_mobile_recomm = ?,
					id_cust_grtee = ?,
					tmp_cust_name_grtee = ?,
					tmp_cust_mobile_grtee = ?,
					id_cust_contact = ?,
					tmp_cust_name_contact = ?,
					tmp_cust_mobile_contact = ?
				WHERE id_tmp_doc_core = ?`
		stmt, err := initializers.DB.Prepare(sqlStmt)
		if err != nil {
			log.Fatalln(err)
			myError := errors.New("2.1 upd : " + err.Error())
			return InstTmpDoccore, myError
		}
		defer stmt.Close()
		_, err = stmt.Exec(InstTmpDoccore.IDCustomer, InstTmpDoccore.TmpCustomerName, InstTmpDoccore.TmpCustomerMobile, InstTmpDoccore.IDCustRecomm, InstTmpDoccore.TmpCustNameRecomm, InstTmpDoccore.TmpCustMobileRecomm, InstTmpDoccore.IDCustGrtee, InstTmpDoccore.TmpCustNameGrtee, InstTmpDoccore.TmpCustMobileGrtee, InstTmpDoccore.IDCustContact, InstTmpDoccore.TmpCustNameContact, InstTmpDoccore.TmpCustMobileContact, InstTmpDoccore.IDTmpDocCore)
		if err != nil {
			log.Fatalln(err)
			myError := errors.New("2.1 upd : " + err.Error())
			return InstTmpDoccore, myError
		}
	} else {
		// new insert
		sqlStmt := `INSERT INTO tmp_doc_core(id_tmp_doc_core, id_doc_type, group_type_1in_2out, id_branch, id_action_name, id_customer, tmp_customer_name, tmp_customer_mobile, id_cust_recomm, tmp_cust_name_recomm, tmp_cust_mobile_recomm, id_cust_grtee, tmp_cust_name_grtee, tmp_cust_mobile_grtee, id_cust_contact, tmp_cust_name_contact, tmp_cust_mobile_contact, month_mm, year_20yy, record_create_time, record_create_by_id, record_create_by_name)
		VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,UNIX_TIMESTAMP(),?,? )`
		stmt, err := initializers.DB.Prepare(sqlStmt)
		if err != nil {
			log.Fatalln(err)
			myError := errors.New("2.1 inst : " + err.Error())
			return InstTmpDoccore, myError
		}
		defer stmt.Close()
		_, err = stmt.Exec(InstTmpDoccore.IDTmpDocCore, InstTmpDoccore.IDDocType, InstTmpDoccore.GroupType1in2out, InstTmpDoccore.IDBranch, InstTmpDoccore.IDActionName, InstTmpDoccore.IDCustomer, InstTmpDoccore.TmpCustomerName, InstTmpDoccore.TmpCustomerMobile, InstTmpDoccore.IDCustRecomm, InstTmpDoccore.TmpCustNameRecomm, InstTmpDoccore.TmpCustMobileRecomm, InstTmpDoccore.IDCustGrtee, InstTmpDoccore.TmpCustNameGrtee, InstTmpDoccore.TmpCustMobileGrtee, InstTmpDoccore.IDCustContact, InstTmpDoccore.TmpCustNameContact, InstTmpDoccore.TmpCustMobileContact, InstTmpDoccore.MonthMm, InstTmpDoccore.Year20yy, InstTmpDoccore.RecordCreateByID, InstTmpDoccore.RecordCreateByName)
		if err != nil {
			log.Fatalln(err)
			myError := errors.New("2.2 inst : " + err.Error())
			return InstTmpDoccore, myError
		}
	}

	tmpDoc, err_chk := initializers.DB.Prepare("SELECT id_tmp_doc_core, id_doc_type, group_type_1in_2out, id_branch, id_action_name, id_customer, tmp_customer_name, tmp_customer_mobile, id_cust_recomm, tmp_cust_name_recomm, tmp_cust_mobile_recomm, id_cust_grtee, tmp_cust_name_grtee, tmp_cust_mobile_grtee, id_cust_contact, tmp_cust_name_contact, tmp_cust_mobile_contact, month_mm, year_20yy, record_create_by_id, record_create_by_name FROM tmp_doc_core WHERE id_tmp_doc_core = ?")
	if err_chk != nil {
		log.Fatal(err_chk.Error())
		myError := errors.New("3 : " + err.Error())
		return InstTmpDoccore, myError
	}
	defer tmpDoc.Close()
	row_core := tmpDoc.QueryRow(id_tmp_doc_core)
	var tmpdoccore modelsStrucTable.TmpDocCore
	err = row_core.Scan(
		&tmpdoccore.IDTmpDocCore,
		&tmpdoccore.IDDocType,
		&tmpdoccore.GroupType1in2out,
		&tmpdoccore.IDBranch,
		&tmpdoccore.IDActionName,
		&tmpdoccore.IDCustomer,
		&tmpdoccore.TmpCustomerName,
		&tmpdoccore.TmpCustomerMobile,
		&tmpdoccore.IDCustRecomm,
		&tmpdoccore.TmpCustNameRecomm,
		&tmpdoccore.TmpCustMobileRecomm,
		&tmpdoccore.IDCustGrtee,
		&tmpdoccore.TmpCustNameGrtee,
		&tmpdoccore.TmpCustMobileGrtee,
		&tmpdoccore.IDCustContact,
		&tmpdoccore.TmpCustNameContact,
		&tmpdoccore.TmpCustMobileContact,
		&tmpdoccore.MonthMm,
		&tmpdoccore.Year20yy,
		&tmpdoccore.RecordCreateByID,
		&tmpdoccore.RecordCreateByName,
	)
	if err != nil {
		myError := errors.New("4 : " + err.Error())
		return tmpdoccore, myError
	}
	return tmpdoccore, nil
}

func Listtmpcore(c *fiber.Ctx) ([]modelsStrucTable.TmpDocCore, error) {
	var curBarch modelsStruc.Cur_branch
	if err := c.BodyParser(&curBarch); err != nil {
		return nil, err
	}
	var Listtmpcore []modelsStrucTable.TmpDocCore
	var err error
	sql_statement := "SELECT id_tmp_doc_core, id_doc_type, group_type_1in_2out, id_branch, id_action_name, id_customer, tmp_customer_name, tmp_customer_mobile, id_cust_recomm, tmp_cust_name_recomm, tmp_cust_mobile_recomm, id_cust_grtee, tmp_cust_name_grtee, tmp_cust_mobile_grtee, id_cust_contact, tmp_cust_name_contact, tmp_cust_mobile_contact, month_mm, year_20yy, record_create_by_id, record_create_by_name FROM tmp_doc_core WHERE id_branch = " + strconv.Itoa(curBarch.Id_branch)

	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStrucTable.TmpDocCore
		err = list.Scan(
			&row.IDTmpDocCore,
			&row.IDDocType,
			&row.GroupType1in2out,
			&row.IDBranch,
			&row.IDActionName,
			&row.IDCustomer,
			&row.TmpCustomerName,
			&row.TmpCustomerMobile,
			&row.IDCustRecomm,
			&row.TmpCustNameRecomm,
			&row.TmpCustMobileRecomm,
			&row.IDCustGrtee,
			&row.TmpCustNameGrtee,
			&row.TmpCustMobileGrtee,
			&row.IDCustContact,
			&row.TmpCustNameContact,
			&row.TmpCustMobileContact,
			&row.MonthMm,
			&row.Year20yy,
			&row.RecordCreateByID,
			&row.RecordCreateByName,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.IDTmpDocCore))
			row.Id_tmp_doc_core_enc = id_enc
		}
		Listtmpcore = append(Listtmpcore, row)
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Listtmpcore, nil
}
