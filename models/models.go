package models

import "database/sql"

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Suspect struct {
	Nomor         int
	Cif1          string
	Kode_uker1    int
	Nama_nasabah1 string
	Tgl1          []uint8
	Cif2          string
	Kode_uker2    int
	Nama_nasabah2 string
	Tgl2          []uint8
}
