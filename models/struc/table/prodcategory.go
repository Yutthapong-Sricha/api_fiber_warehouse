package models

type ProdCategory struct {
	IDProdCategory       int     `json:"id_prod_category" db:"id_prod_category"`           // id_prod_category
	CategoryName         *string `json:"category_name" db:"category_name"`                 // category_name
	IsActiveFlag         *int    `json:"is_active_flag" db:"is_active_flag"`               // is_active_flag
	RecordCreateTime     *int    `json:"record_create_time" db:"record_create_time"`       // record_create_time
	RecordCreateByID     *int    `json:"record_create_by_id" db:"record_create_by_id"`     // record_create_by_id
	RecordCreateByName   *string `json:"record_create_by_name" db:"record_create_by_name"` // record_create_by_name
	RecordUpdateTime     *int    `json:"record_update_time" db:"record_update_time"`       // record_update_time
	RecordUpdateByID     *int    `json:"record_update_by_id" db:"record_update_by_id"`     // record_update_by_id
	RecordUpdateByName   *string `json:"record_update_by_name" db:"record_update_by_name"` // record_update_by_name
	CategoryIDOld        *int    `json:"CategoryId_old" db:"CategoryId_old"`               // CategoryId_old
	Id_prod_category_enc string  `json:"id_prod_category_enc"`
}
