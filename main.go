package main

import (
	"net/http"

	"github.com/jeypc/go-crud-modal/controllers/mahasiswacontroller"
)

func main() {
	http.HandleFunc("/", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/get_form", mahasiswacontroller.GetForm)
	http.HandleFunc("/mahasiswa/store", mahasiswacontroller.Store)
	http.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)

	http.ListenAndServe(":8000", nil)
}
