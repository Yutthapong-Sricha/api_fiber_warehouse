package models

type AddrProvince struct {
	IDAddrProvince    *int    `json:"id_addr_province" db:"id_addr_province"`       // id_addr_province
	ProvNameTh        *string `json:"prov_name_th" db:"prov_name_th"`               // prov_name_th
	ProvNameEn        *string `json:"prov_name_en" db:"prov_name_en"`               // prov_name_en
	IDAddrGeographies *int    `json:"id_addr_geographies" db:"id_addr_geographies"` // id_addr_geographies
	CreateTime        *string `json:"create_time" db:"create_time"`                 // create_time
	UpdateTime        *string `json:"update_time" db:"update_time"`                 // update_time
	DeletedTime       *string `json:"deleted_time" db:"deleted_time"`               // deleted_time
}
