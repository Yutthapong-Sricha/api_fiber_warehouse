package models

type AddrTambon struct {
	IDAddrTambon  *int    `json:"id_addr_tambon" db:"id_addr_tambon"`   // id_addr_tambon
	ZipCode       *int    `json:"zip_code"    db:"zip_code"`            // zip_code
	TamNameTh     *string `json:"tam_name_th"   db:"tam_name_th"`       // tam_name_th
	TamNameEn     *string `json:"tam_name_en"   db:"tam_name_en"`       // tam_name_en
	IDAddrAmphure *int    `json:"id_addr_amphure" db:"id_addr_amphure"` // id_addr_amphure
	CreatedTime   *string `json:"created_time"  db:"created_time"`      // created_time
	UpdatedTime   *string `json:"updated_time"  db:"updated_time"`      // updated_time
	DeletedTime   *string `json:"deleted_time"  db:"deleted_time"`      // deleted_time
}
