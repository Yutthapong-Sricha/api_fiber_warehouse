package models

type Custtmpdoccore struct {
	Id_tmp_doc_core         int    `json:"id_tmp_doc_core" db:"id_tmp_doc_core"`
	Id_doc_type             int    `json:"id_doc_type" db:"id_doc_type"`
	Group_type_1in_2out     int    `json:"group_type_1in_2out" db:"group_type_1in_2out"`
	Id_branch               int    `json:"id_branch" db:"id_branch"`
	Id_action_name          int    `json:"id_action_name" db:"id_action_name"`
	Id_customer             int    `json:"id_customer" db:"id_customer"`
	Tmp_customer_name       string `json:"tmp_customer_name" db:"tmp_customer_name"`
	Tmp_customer_mobile     string `json:"tmp_customer_mobile" db:"tmp_customer_mobile"`
	Id_cust_recomm          int    `json:"id_cust_recomm" db:"id_cust_recomm"`
	Tmp_cust_name_recomm    string `json:"tmp_cust_name_recomm" db:"tmp_cust_name_recomm"`
	Tmp_cust_mobile_recomm  string `json:"tmp_cust_mobile_recomm" db:"tmp_cust_mobile_recomm"`
	Id_cust_grtee           int    `json:"id_cust_grtee" db:"id_cust_grtee"`
	Tmp_cust_name_grtee     string `json:"tmp_cust_name_grtee" db:"tmp_cust_name_grtee"`
	Tmp_cust_mobile_grtee   string `json:"tmp_cust_mobile_grtee" db:"tmp_cust_mobile_grtee"`
	Id_cust_contact         int    `json:"id_cust_contact" db:"id_cust_contact"`
	Tmp_cust_name_contact   string `json:"tmp_cust_name_contact" db:"tmp_cust_name_contact"`
	Tmp_cust_mobile_contact string `json:"tmp_cust_mobile_contact" db:"tmp_cust_mobile_contact"`
	Month_mm                string `json:"month_mm" db:"month_mm"`
	Year_20yy               string `json:"year_20yy" db:"year_20yy"`
	Record_create_by_id     int    `json:"record_create_by_id" db:"record_create_by_id"`
	Record_create_by_name   string `json:"record_create_by_name" db:"record_create_by_name"`
}

type Tmpdoccore struct {
	Id_tmp_doc_core         int    `json:"id_tmp_doc_core" db:"id_tmp_doc_core"`
	Id_doc_type             int    `json:"id_doc_type" db:"id_doc_type"`
	Group_type_1in_2out     int    `json:"group_type_1in_2out" db:"group_type_1in_2out"`
	Id_branch               int    `json:"id_branch" db:"id_branch"`
	Id_action_name          int    `json:"id_action_name" db:"id_action_name"`
	Id_customer             int    `json:"id_customer" db:"id_customer"`
	Tmp_customer_name       string `json:"tmp_customer_name" db:"tmp_customer_name"`
	Tmp_customer_mobile     string `json:"tmp_customer_mobile" db:"tmp_customer_mobile"`
	Id_cust_recomm          int    `json:"id_cust_recomm" db:"id_cust_recomm"`
	Tmp_cust_name_recomm    string `json:"tmp_cust_name_recomm" db:"tmp_cust_name_recomm"`
	Tmp_cust_mobile_recomm  string `json:"tmp_cust_mobile_recomm" db:"tmp_cust_mobile_recomm"`
	Id_cust_grtee           int    `json:"id_cust_grtee" db:"id_cust_grtee"`
	Tmp_cust_name_grtee     string `json:"tmp_cust_name_grtee" db:"tmp_cust_name_grtee"`
	Tmp_cust_mobile_grtee   string `json:"tmp_cust_mobile_grtee" db:"tmp_cust_mobile_grtee"`
	Id_cust_contact         int    `json:"id_cust_contact" db:"id_cust_contact"`
	Tmp_cust_name_contact   string `json:"tmp_cust_name_contact" db:"tmp_cust_name_contact"`
	Tmp_cust_mobile_contact string `json:"tmp_cust_mobile_contact" db:"tmp_cust_mobile_contact"`
	Month_mm                string `json:"month_mm" db:"month_mm"`
	Year_20yy               string `json:"year_20yy" db:"year_20yy"`
	Record_create_by_id     int    `json:"record_create_by_id" db:"record_create_by_id"`
	Record_create_by_name   string `json:"record_create_by_name" db:"record_create_by_name"`
	Id_tmp_doc_core_enc     string `json:"id_tmp_doc_core_enc" db:"id_tmp_doc_core_enc"`
}
