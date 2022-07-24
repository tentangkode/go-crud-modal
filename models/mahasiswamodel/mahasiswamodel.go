package mahasiswamodel

import (
	"database/sql"

	"github.com/jeypc/go-crud-modal/config"
	"github.com/jeypc/go-crud-modal/entities"
)

type MahasiswaModel struct {
	db *sql.DB
}

func New() *MahasiswaModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &MahasiswaModel{db: db}
}

func (m *MahasiswaModel) FindAll(mahasiswa *[]entities.Mahasiswa) error {
	rows, err := m.db.Query("select * from mahasiswa")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Mahasiswa
		rows.Scan(
			&data.Id,
			&data.NamaLengkap,
			&data.JenisKelamin,
			&data.TempatLahir,
			&data.TanggalLahir,
			&data.Alamat)
		*mahasiswa = append(*mahasiswa, data)
	}

	return nil
}

func (m *MahasiswaModel) Create(mahasiswa *entities.Mahasiswa) error {
	result, err := m.db.Exec("insert into mahasiswa (nama_lengkap, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat) values(?,?,?,?,?)",
		mahasiswa.NamaLengkap, mahasiswa.JenisKelamin, mahasiswa.TempatLahir, mahasiswa.TanggalLahir, mahasiswa.Alamat)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	mahasiswa.Id = lastInsertId
	return nil
}

func (m *MahasiswaModel) Find(id int64, mahasiswa *entities.Mahasiswa) error {
	return m.db.QueryRow("select * from mahasiswa where id = ?", id).Scan(
		&mahasiswa.Id,
		&mahasiswa.NamaLengkap,
		&mahasiswa.JenisKelamin,
		&mahasiswa.TempatLahir,
		&mahasiswa.TanggalLahir,
		&mahasiswa.Alamat)
}

func (m *MahasiswaModel) Update(mahasiswa entities.Mahasiswa) error {

	_, err := m.db.Exec("update mahasiswa set nama_lengkap = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ? where id = ?",
		mahasiswa.NamaLengkap, mahasiswa.JenisKelamin, mahasiswa.TempatLahir, mahasiswa.TanggalLahir, mahasiswa.Alamat, mahasiswa.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *MahasiswaModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from mahasiswa where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
