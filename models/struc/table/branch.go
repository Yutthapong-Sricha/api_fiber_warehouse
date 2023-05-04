package models

type Branch struct {
	IDBranch           int     `json:"id_branch" db:"id_branch"`
	BranchName         *string `json:"branch_name" db:"branch_name"`
	Address            *string `json:"address" db:"address"`
	IDAddrTambon       *int    `json:"id_addr_tambon" db:"id_addr_tambon"`
	IDAddrAmphure      *int    `json:"id_addr_amphure" db:"id_addr_amphure"`
	IDAddrProvince     *int    `json:"id_addr_province" db:"id_addr_province"`
	BranchTel          *string `json:"branch_tel" db:"branch_tel"`
	BranchMobile       *string `json:"branch_mobile" db:"branch_mobile"`
	BranchLineID       *string `json:"branch_line_id" db:"branch_line_id"`
	IsActiveFlag       *int    `json:"is_active_flag" db:"is_active_flag"`
	RecordUpdateTime   *int    `json:"record_update_time" db:"record_update_time"`
	RecordUpdateByID   *int    `json:"record_update_by_id" db:"record_update_by_id"`
	RecordUpdateByName *string `json:"record_update_by_name" db:"record_update_by_name"`
	Id_branch_enc      string  `json:"id_branch_enc" db:"id_branch_enc"`
}
