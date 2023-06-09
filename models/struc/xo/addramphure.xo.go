package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// AddrAmphure represents a row from 'spc_holding.addr_amphure'.
type AddrAmphure struct {
	IDAddrAmphure  sql.NullInt64  `json:"id_addr_amphure"`  // id_addr_amphure
	AmpNameTh      sql.NullString `json:"amp_name_th"`      // amp_name_th
	AmpNameEn      sql.NullString `json:"amp_name_en"`      // amp_name_en
	IDAddrProvince sql.NullInt64  `json:"id_addr_province"` // id_addr_province
	CreatedTime    sql.NullString `json:"created_time"`     // created_time
	UpdatedTime    sql.NullString `json:"updated_time"`     // updated_time
	DeletedTime    sql.NullString `json:"deleted_time"`     // deleted_time
}
