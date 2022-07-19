package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOne(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	nomor, err := strconv.Atoi(params.ByName("nomor"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJson(w, err)
		return
	}
	// app.logger.Println("nomor is", nomor)
	// suspect := models.Suspect{
	// 	Nomor:         nomor,
	// 	Cif1:          "tes",
	// 	Kode_uker1:    3,
	// 	Nama_nasabah1: "tes",
	// 	Tgl1:          20,
	// 	Cif2:          "tes",
	// 	Kode_uker2:    34,
	// 	Nama_nasabah2: "tes",
	// 	Tgl2:          20,

	app.logger.Println("nomor is", nomor)

	suspect, err := app.models.DB.Get(nomor)

	err = app.writeJson(w, http.StatusOK, suspect, "suspect")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}
func (app *application) getAll(w http.ResponseWriter, r *http.Request) {
	suspects, err := app.models.DB.All()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, suspects, "suspects")
	if err != nil {
		app.errorJson(w, err)
		return
	}
}
