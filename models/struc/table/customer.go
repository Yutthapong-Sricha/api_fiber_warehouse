package models

type Customer struct {
	IDCustomer         int     `json:"id_customer" db:"id_customer"`                     //id_customer
	CustCodeSpc        *string `json:"cust_code_spc" db:"cust_code_spc"`                 //cust_code_spc
	CustName           *string `json:"cust_name" db:"cust_name"`                         //cust_name
	CustLastname       *string `json:"cust_lastname" db:"cust_lastname"`                 //cust_lastname
	CustFullname       *string `json:"cust_fullname" db:"cust_fullname"`                 //cust_fullname
	CustNickname       *string `json:"cust_nickname" db:"cust_nickname"`                 //cust_nickname
	CustTaxID          *string `json:"cust_tax_id" db:"cust_tax_id"`                     //cust_tax_id
	CustMobile         *string `json:"cust_mobile" db:"cust_mobile"`                     //cust_mobile
	CustTel            *string `json:"cust_tel" db:"cust_tel"`                           //cust_tel
	CustEmail          *string `json:"cust_email" db:"cust_email"`                       //cust_email
	CustAddress        *string `json:"cust_address" db:"cust_address"`                   //cust_address
	IDAddrTambon       *int    `json:"id_addr_tambon" db:"id_addr_tambon"`               //id_addr_tambon
	TmpTambonName      *string `json:"tmp_tambon_name" db:"tmp_tambon_name"`             //tmp_tambon_name
	IDAddrAmphure      *int    `json:"id_addr_amphure" db:"id_addr_amphure"`             //id_addr_amphure
	TmpAmphureName     *string `json:"tmp_amphure_name" db:"tmp_amphure_name"`           //tmp_amphure_name
	IDAddrProvince     *int    `json:"id_addr_province" db:"id_addr_province"`           //id_addr_province
	TmpProvinceName    *string `json:"tmp_province_name" db:"tmp_province_name"`         //tmp_province_name
	TmpZipCode         *string `json:"tmp_zip_code" db:"tmp_zip_code"`                   //tmp_zip_code
	Pers4OrComp8       *int    `json:"pers4_or_comp8" db:"pers4_or_comp8"`               //pers4_or_comp8
	ContactName        *string `json:"contact_name" db:"contact_name"`                   //contact_name
	ContactNote        *string `json:"contact_note" db:"contact_note"`                   //contact_note
	IsMember356        *int    `json:"is_member_356" db:"is_member_356"`                 //is_member_356
	LineID             *string `json:"line_id" db:"line_id"`                             //line_id
	Cust356Flag        *int    `json:"cust_356_flag" db:"cust_356_flag"`                 //cust_356_flag
	RecordCreateTime   *int    `json:"record_create_time" db:"record_create_time"`       //record_create_time
	RecordCreateByID   *int    `json:"record_create_by_id" db:"record_create_by_id"`     //record_create_by_id
	RecordCreateByName *string `json:"record_create_by_name" db:"record_create_by_name"` //record_create_by_name
	RecordUpdateTime   *int    `json:"record_update_time" db:"record_update_time"`       //record_update_time
	RecordUpdateByID   *int    `json:"record_update_by_id" db:"record_update_by_id"`     //record_update_by_id
	RecordUpdateByName *string `json:"record_update_by_name" db:"record_update_by_name"` //record_update_by_name
	Note               *string `json:"note" db:"note"`                                   //note
	IDCustomer_enc     string  `json:"id_customer_enc" db:"id_customer_enc"`
}
