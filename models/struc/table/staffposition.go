package models

type StaffPosition struct {
	IDStaffPosition    int     `json:"id_staff_position" db:"id_staff_position"`
	PositionName       *string `json:"position_name" db:"position_name"`
	RowOrderPosition   *int    `json:"row_order_position" db:"row_order_position"`
	IsActiveFlag       *int    `json:"is_active_flag" db:"is_active_flag"`
	RecordUpdateTime   *int    `json:"record_update_time" db:"record_update_time"`
	RecordUpdateByID   *int    `json:"record_update_by_id" db:"record_update_by_id"`
	RecordUpdateByName *string `json:"record_update_by_name" db:"record_update_by_name"`
	Id_position_enc    string  `json:"id_position_enc" db:"id_position_enc"`
}
