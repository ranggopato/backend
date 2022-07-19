package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one suspect and error, if any
func (m *DBModel) Get(nomor int) (*Suspect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select nomor, cif1, kode_uker1, nama_nasabah1, tgl1, cif2, kode_uker2, nama_nasabah2,tgl2 from suspect_mark where nomor = $1`

	row := m.DB.QueryRowContext(ctx, query, nomor)

	var suspect Suspect

	err := row.Scan(
		&suspect.Nomor,
		&suspect.Cif1,
		&suspect.Kode_uker1,
		&suspect.Nama_nasabah1,
		&suspect.Tgl1,
		&suspect.Cif2,
		&suspect.Kode_uker2,
		&suspect.Nama_nasabah2,
		&suspect.Tgl2,
	)
	if err != nil {
		return nil, err
	}
	return &suspect, nil
}

// All returns all suspects and error, if any
func (m *DBModel) All() ([]*Suspect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select nomor, cif1, kode_uker1, nama_nasabah1, tgl1, cif2, kode_uker2, nama_nasabah2, tgl2 from suspect_mark"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suspects []*Suspect

	for rows.Next() {
		var suspect Suspect
		err := rows.Scan(
			&suspect.Nomor,
			&suspect.Cif1,
			&suspect.Kode_uker1,
			&suspect.Nama_nasabah1,
			&suspect.Tgl1,
			&suspect.Cif2,
			&suspect.Kode_uker2,
			&suspect.Nama_nasabah2,
			&suspect.Tgl2,
		)
		if err != nil {
			return nil, err
		}

		suspects = append(suspects, &suspect)

	}
	return suspects, nil
}
