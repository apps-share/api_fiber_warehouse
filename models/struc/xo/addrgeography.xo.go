package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// AddrGeography represents a row from 'spc_holding.addr_geographies'.
type AddrGeography struct {
	IDAddrGeographies sql.NullInt64  `json:"id_addr_geographies"` // id_addr_geographies
	GeoName           sql.NullString `json:"geo_name"`            // geo_name
}
