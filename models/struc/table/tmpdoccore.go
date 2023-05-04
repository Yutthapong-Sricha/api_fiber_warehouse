package models

// TmpDocCore represents a row from 'spc_holding.tmp_doc_core'.
type TmpDocCore struct {
	IDTmpDocCore         int      `json:"id_tmp_doc_core"  db:"id_tmp_doc_core"`                // id_tmp_doc_core
	IDDocType            *int     `json:"id_doc_type"   db:"id_doc_type"`                       // id_doc_type
	GroupType1in2out     int      `json:"group_type_1in_2out" db:"group_type_1in_2out"`         // group_type_1in_2out
	IDBranch             *int     `json:"id_branch"    db:"id_branch"`                          // id_branch
	IDActionName         *int     `json:"id_action_name"   db:"id_action_name"`                 // id_action_name
	DocNumber            *string  `json:"doc_number"    db:"doc_number"`                        // doc_number
	DocRunning           *int     `json:"doc_running"   db:"doc_running"`                       // doc_running
	IDCustomer           *int     `json:"id_customer"   db:"id_customer"`                       // id_customer
	TmpCustomerName      *string  `json:"tmp_customer_name"  db:"tmp_customer_name"`            // tmp_customer_name
	TmpCustomerMobile    *string  `json:"tmp_customer_mobile" db:"tmp_customer_mobile"`         // tmp_customer_mobile
	IDCustRecomm         *int     `json:"id_cust_recomm"   db:"id_cust_recomm"`                 // id_cust_recomm
	TmpCustNameRecomm    *string  `json:"tmp_cust_name_recomm" db:"tmp_cust_name_recomm"`       // tmp_cust_name_recomm
	TmpCustMobileRecomm  *string  `json:"tmp_cust_mobile_recomm" db:"tmp_cust_mobile_recomm"`   // tmp_cust_mobile_recomm
	IDCustParent         *int     `json:"id_cust_parent"   db:"id_cust_parent"`                 // id_cust_parent
	TmpCustNameParent    *string  `json:"tmp_cust_name_parent" db:"tmp_cust_name_parent"`       // tmp_cust_name_parent
	TmpCustMobileParent  *string  `json:"tmp_cust_mobile_parent" db:"tmp_cust_mobile_parent"`   // tmp_cust_mobile_parent
	IDCustGrtee          *int     `json:"id_cust_grtee"   db:"id_cust_grtee"`                   // id_cust_grtee
	TmpCustNameGrtee     *string  `json:"tmp_cust_name_grtee" db:"tmp_cust_name_grtee"`         // tmp_cust_name_grtee
	TmpCustMobileGrtee   *string  `json:"tmp_cust_mobile_grtee" db:"tmp_cust_mobile_grtee"`     // tmp_cust_mobile_grtee
	IDCustContact        *int     `json:"id_cust_contact"  db:"id_cust_contact"`                // id_cust_contact
	TmpCustNameContact   *string  `json:"tmp_cust_name_contact" db:"tmp_cust_name_contact"`     // tmp_cust_name_contact
	TmpCustMobileContact *string  `json:"tmp_cust_mobile_contact" db:"tmp_cust_mobile_contact"` // tmp_cust_mobile_contact
	IDSupplier           *int     `json:"id_supplier"   db:"id_supplier"`                       // id_supplier
	TmpSuppName          *string  `json:"tmp_supp_name"   db:"tmp_supp_name"`                   // tmp_supp_name
	TmpContact1Mobile    *string  `json:"tmp_contact1_mobile" db:"tmp_contact1_mobile"`         // tmp_contact1_mobile
	TmpContact1Email     *string  `json:"tmp_contact1_email"  db:"tmp_contact1_email"`          // tmp_contact1_email
	TotalQuantity        *int     `json:"total_quantity"   db:"total_quantity"`                 // total_quantity
	TotalPrice           *float64 `json:"total_price"   db:"total_price"`                       // total_price
	TotalDownPayment     *float64 `json:"total_down_payment"  db:"total_down_payment"`          // total_down_payment
	TotalAfterPayDown    *float64 `json:"total_after_pay_down" db:"total_after_pay_down"`       // total_after_pay_down
	TotalDiscountManual  *float64 `json:"total_discount_manual" db:"total_discount_manual"`     // total_discount_manual
	TotalCalVat          *float64 `json:"total_cal_vat"   db:"total_cal_vat"`                   // total_cal_vat
	TotalNetPrice        *float64 `json:"total_net_price"  db:"total_net_price"`                // total_net_price
	DocCoreDesc          *string  `json:"doc_core_desc"   db:"doc_core_desc"`                   // doc_core_desc
	RecordCreateTime     *int     `json:"record_create_time"  db:"record_create_time"`          // record_create_time
	RecordCreateByID     *int     `json:"record_create_by_id" db:"record_create_by_id"`         // record_create_by_id
	RecordCreateByName   *string  `json:"record_create_by_name" db:"record_create_by_name"`     // record_create_by_name
	IDSaleStaff          *int     `json:"id_sale_staff"   db:"id_sale_staff"`                   // id_sale_staff
	TmpSaleStaffName     *int     `json:"tmp_sale_staff_name" db:"tmp_sale_staff_name"`         // tmp_sale_staff_name
	MonthMm              string   `json:"month_mm"    db:"month_mm"`                            // month_mm
	Year20yy             string   `json:"year_20yy"    db:"year_20yy"`                          // year_20yy
	TotalTrans           *int     `json:"total_trans"   db:"total_trans"`                       // total_trans
	HasEditTrans356      *int     `json:"has_edit_trans_356"  db:"has_edit_trans_356"`          // has_edit_trans_356
	TotalEditTrans       *int     `json:"total_edit_trans"  db:"total_edit_trans"`              // total_edit_trans
	SplitDownPaymentFlag *int     `json:"split_down_payment_flag" db:"split_down_payment_flag"` //
	Id_tmp_doc_core_enc  string   `json:"id_tmp_doc_core_enc"`
}
