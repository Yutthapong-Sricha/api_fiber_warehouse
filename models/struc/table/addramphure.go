package models

type AddrAmphure struct {
	IDAddrAmphure  *int    `json:"id_addr_amphure"  db:"id_addr_amphure"`   // id_addr_amphure
	AmpNameTh      *string `json:"amp_name_th"  db:"amp_name_th"`           // amp_name_th
	AmpNameEn      *string `json:"amp_name_en"  db:"amp_name_en"`           // amp_name_en
	IDAddrProvince *int    `json:"id_addr_province"  db:"id_addr_province"` // id_addr_province
	CreatedTime    *string `json:"created_time"  db:"created_time"`         // created_time
	UpdatedTime    *string `json:"updated_time"  db:"updated_time"`         // updated_time
	DeletedTime    *string `json:"deleted_time"  db:"deleted_time"`         // deleted_time
}
